package chat

import (
	"backend/pkg/models"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	ai "backend/pkg/services/ai/llm"
	"backend/pkg/services/templates"
)

// ChatService manages both AI-driven and user-user chat interactions.
type ChatService struct {
	geminiService   *ai.GeminiService
	templateService *templates.TemplateService
	conversationDB  ConversationDB
	convCache       map[string]*Conversation // In-memory cache.
	mu              sync.Mutex
}

// ConversationDB interface to define methods for interacting with the DB.
type ConversationDB interface {
	GetConversation(convID string) (*Conversation, error)
	SaveConversation(conv *Conversation) error
	DataBuilder(chatType string, outputLength int32) map[string]interface{}
}

// Conversation holds in-memory chat data.
type Conversation struct {
	ID               string
	FullConversation []models.Dialogue
	RecentDialogue   []models.Dialogue
	Summary          string
	IsActive         bool
}

// GeminiService returns the geminiService instance.
func (cs *ChatService) GeminiService() *ai.GeminiService {
	return cs.geminiService
}

// InitializeChatService initializes the ChatService.
func InitializeChatService(db ConversationDB, templateService *templates.TemplateService) (*ChatService, error) {
	gemini, err := ai.InitializeGeminiService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GeminiService: %w", err)
	}
	return &ChatService{
		geminiService:   gemini,
		templateService: templateService,
		conversationDB:  db,
		convCache:       make(map[string]*Conversation),
	}, nil
}

// EndChat ends the conversation session and saves it to the database.
func (cs *ChatService) EndChat(convID string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conversation, err := cs.getConversationFromCache(convID)
	if err != nil {
		return err
	}

	// Save the full conversation to the DB.
	err = cs.conversationDB.SaveConversation(conversation)
	if err != nil {
		return fmt.Errorf("failed to save conversation: %w", err)
	}

	conversation.IsActive = false
	cs.geminiService.EndChat()
	return nil
}

// PredictChat adds a user message, generates a response from the AI, and updates the conversation.
func (cs *ChatService) PredictChat(ctx context.Context, convID, chatType, userMessage string) (string, error) {
	cs.mu.Lock()
	conversation, err := cs.getConversationFromCache(convID)
	cs.mu.Unlock()

	if err != nil {
		return "", err
	}

	// Generate response using GeminiService.
	response, err := cs.geminiService.PredictChat(ctx, userMessage)
	if err != nil {
		return "", fmt.Errorf("failed to predict response: %w", err)
	}

	// Update conversation history.
	cs.updateConversationHistory(conversation, models.Dialogue{Role: models.RoleUser, Content: userMessage}, models.Dialogue{Role: models.RoleModel, Content: response})

	// Check if conversation exceeds character limit for summarization.
	if calculateContextLength(conversation.FullConversation) > MaxConversationLength {
		summary, err := cs.generateSummary(ctx, conversation.FullConversation)
		if err != nil {
			return "", fmt.Errorf("failed to generate summary: %w", err)
		}
		cs.restartConversationWithSummary(conversation, summary, chatType)
	}

	return response, nil
}

// StartChat initializes a conversation session based on the specified chat type.
func (cs *ChatService) StartChat(convID string, chatType string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	maxTokens, temperature := setParametersByChatType(chatType)
	log.Printf("Starting chat with type '%s' and max tokens '%d'", chatType, maxTokens)

	templateName, err := getTemplateNameByChatType(chatType)
	if err != nil {
		return err
	}

	data := cs.conversationDB.DataBuilder(chatType, maxTokens)
	if data == nil {
		return errors.New("failed to build prompt data")
	}
	log.Printf("Prompt data: %v", data)

	// Render the prompt using the template service.
	prompt, err := cs.templateService.RenderTemplate(templateName, data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}
	log.Printf("Prompt: %s", prompt)
	// Load conversation from cache or DB.
	conversation, err := cs.getOrLoadConversation(convID)
	if err != nil {
		return fmt.Errorf("failed to load conversation: %w", err)
	}
	log.Printf("Conversation loaded: %v", conversation)
	// Ensure that the chat type matches the existing conversation type (if applicable).
	if conversation.IsActive {
		return errors.New("conversation already in progress")
	}
	log.Printf("Starting new chat session with prompt: %s", prompt)
	// Start new chat session with system instruction based on the prompt.
	err = cs.geminiService.StartNewChat(prompt, conversation.RecentDialogue, maxTokens, temperature)
	if err != nil {
		return fmt.Errorf("failed to start new chat: %w", err)
	}
	conversation.IsActive = true
	return nil
}

// getConversationFromCache retrieves a conversation from the cache with proper error handling.
func (cs *ChatService) getConversationFromCache(convID string) (*Conversation, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conversation, ok := cs.convCache[convID]
	if !ok {
		return nil, errors.New("conversation not found")
	}
	return conversation, nil
}

// getOrLoadConversation loads a conversation from the cache or DB if not present.
func (cs *ChatService) getOrLoadConversation(convID string) (*Conversation, error) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conversation, ok := cs.convCache[convID]
	if !ok {
		conv, err := cs.conversationDB.GetConversation(convID)
		if err != nil {
			return nil, fmt.Errorf("failed to load conversation from DB: %w", err)
		}
		cs.convCache[convID] = conv
		conversation = conv
	}
	return conversation, nil
}

// updateConversationHistory updates the conversation history in-memory.
func (cs *ChatService) updateConversationHistory(conversation *Conversation, userDialogue, modelDialogue models.Dialogue) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conversation.FullConversation = append(conversation.FullConversation, userDialogue, modelDialogue)

	// Maintain recent dialogues for summarization purposes (last 6 dialogues).
	conversation.RecentDialogue = append(conversation.RecentDialogue, userDialogue, modelDialogue)
	if len(conversation.RecentDialogue) > 6 {
		conversation.RecentDialogue = conversation.RecentDialogue[len(conversation.RecentDialogue)-6:]
	}
}

// generateSummary uses GeminiService to generate a summary of the conversation.
func (cs *ChatService) generateSummary(ctx context.Context, dialogues []models.Dialogue) (string, error) {
	maxTokens, temperature := setParametersByChatType("summarization")
	prompt, err := cs.generateSummarizationPrompt(dialogues, maxTokens)
	if err != nil {
		return "", err
	}
	return cs.geminiService.Predict(ctx, prompt, maxTokens, temperature)
}

// restartConversationWithSummary restarts the conversation after generating a summary.
func (cs *ChatService) restartConversationWithSummary(conversation *Conversation, summary, chatType string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	conversation.Summary = summary
	maxTokens, temperature := setParametersByChatType(chatType)
	err := cs.geminiService.StartNewChat(summary, conversation.RecentDialogue, maxTokens, temperature)
	if err != nil {
		log.Printf("Failed to restart conversation with summary: %v", err)
	}
}

// generateSummarizationPrompt builds the prompt for summarization.
func (cs *ChatService) generateSummarizationPrompt(dialogues []models.Dialogue, outputLength int32) (string, error) {
	var sb strings.Builder
	for _, d := range dialogues {
		sb.WriteString(fmt.Sprintf("%s: %s\n", d.Role, d.Content))
	}

	data := map[string]interface{}{
		"conversation":  sb.String(),
		"output_length": outputLength,
	}
	summary, err := cs.templateService.RenderTemplate(templates.SUMMARIZATION_MODE, data)
	if err != nil {
		return "", fmt.Errorf("failed to render summarization template: %w", err)
	}
	return summary, nil
}
