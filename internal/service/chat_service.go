package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"sync"
)

type ChatService struct {
	geminiService   *GeminiService
	templateService *TemplateService
	conversationDB  *DBConversationStore
}

// InMemoryChatStore stores active conversations in-memory for batching
// Recent dialogues should be added here, not in DB
type InMemoryChatStore struct {
	activeConversations map[string]*entity.Conversation
	recentDialogues     entity.Dialogue
	mu                  sync.Mutex
}

func NewInMemoryChatStore() *InMemoryChatStore {
	return &InMemoryChatStore{
		activeConversations: make(map[string]*entity.Conversation),
		recentDialogues:     entity.Dialogue{},
	}
}

func getRecentDialogues() entity.Dialogue {
	// Get 6 recent dialogues from the active conversation just loaded
	// from the database
	return entity.Dialogue{}
}

// InitializeChatService initializes the ChatService with the database store
func InitializeChatService(db *DBConversationStore, templateService *TemplateService) (*ChatService, error) {
	gemini, err := InitializeGeminiService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize GeminiService: %w", err)
	}
	return &ChatService{
		geminiService:   gemini,
		templateService: templateService,
		conversationDB:  db,
	}, nil
}

// EndChat ends the conversation session and saves it to the database.
func (cs *ChatService) EndChat(convID string) error {
	conversation, err := cs.getConversation(convID)
	if err != nil {
		return err
	}

	conversation.IsActive = false

	// Save the conversation to the database.
	err = cs.conversationDB.SaveConversation(conversation)
	if err != nil {
		return fmt.Errorf("failed to save conversation: %w", err)
	}

	cs.geminiService.EndChat()
	return nil
}

// PredictChat adds a user message, generates a response from the AI, and updates the conversation.
func (cs *ChatService) PredictChat(ctx context.Context, convID, chatType, userMessage string) (string, error) {
	conversation, err := cs.getConversation(convID)
	if err != nil {
		return "", err
	}

	// Generate response using GeminiService.
	response, err := cs.geminiService.PredictChat(ctx, userMessage)
	if err != nil {
		return "", fmt.Errorf("failed to predict response: %w", err)
	}

	// Update conversation history.
	cs.updateConversationHistory(conversation, entity.Dialogue{Role: entity.RoleUser, Content: userMessage}, entity.Dialogue{Role: entity.RoleModel, Content: response})

	// Save updated conversation.
	err = cs.conversationDB.SaveConversation(conversation)
	if err != nil {
		return "", fmt.Errorf("failed to save conversation: %w", err)
	}

	return response, nil
}

// StartChat initializes a conversation session based on the specified chat type.
func (cs *ChatService) StartChat(convID string, chatType string) error {
	maxTokens, temperature := setParametersByChatType(chatType)

	templateName, err := getTemplateNameByChatType(chatType)
	if err != nil {
		return err
	}

	data := cs.conversationDB.DataBuilder(chatType, maxTokens)

	// Render the prompt using the template service.
	prompt, err := cs.templateService.RenderTemplate(templateName, data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	conversation, err := cs.getConversation(convID)
	if err != nil {
		return fmt.Errorf("failed to load conversation: %w", err)
	}

	if conversation.IsActive {
		return errors.New("conversation already in progress")
	}

	err = cs.geminiService.StartNewChat(prompt, conversation.RecentDialogues, maxTokens, temperature)
	if err != nil {
		return fmt.Errorf("failed to start new chat: %w", err)
	}

	conversation.IsActive = true

	err = cs.conversationDB.SaveConversation(conversation)
	if err != nil {
		return fmt.Errorf("failed to save conversation: %w", err)
	}

	return nil
}

// getConversation retrieves a conversation from the database.
func (cs *ChatService) getConversation(convID string) (*entity.Conversation, error) {
	conversation, err := cs.conversationDB.GetConversation(convID)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// updateConversationHistory updates the conversation history.
func (cs *ChatService) updateConversationHistory(conversation *entity.Conversation, userDialogue, modelDialogue entity.Dialogue) {
	conversation.Conversation = append(conversation.Conversation, userDialogue, modelDialogue)

	// Maintain recent dialogues for summarization purposes (last 6 dialogues).
	conversation.RecentDialogues = append(conversation.RecentDialogues, userDialogue, modelDialogue)
	if len(conversation.RecentDialogues) > 6 {
		conversation.RecentDialogues = conversation.RecentDialogues[len(conversation.RecentDialogues)-6:]
	}
}

func (cs *ChatService) CloseGeminiService() {
	cs.geminiService.Close()
}

// getTemplateNameByChatType returns the template name based on the chat type.
func getTemplateNameByChatType(chatType string) (string, error) {
	switch chatType {
	case "recipe":
		return RECIPE_MODE, nil
	case "shopping":
		return SHOPPING_MODE, nil
	case "health":
		return HEALTH_MODE, nil
	case "motivation":
		return MOTIVATION_MODE, nil
	case "calorie_tracker":
		return CALORIE_TRACKER_MODE, nil
	case "summarization":
		return CHAT_SUMMARIZATION_MODE, nil
	default:
		return "", errors.New("invalid chat type")
	}
}

func setParametersByChatType(chatType string) (int32, float32) {
	switch chatType {
	case "recipe":
		return setResponseLength("long"), setTemperature("creative")
	case "shopping":
		return setResponseLength("short"), setTemperature("regular")
	case "health":
		return setResponseLength("long"), setTemperature("regular")
	case "motivation":
		return setResponseLength("medium"), setTemperature("creative")
	case "calorie_tracker":
		return setResponseLength("medium"), setTemperature("deterministic")
	case "summarization":
		return setResponseLength("long"), setTemperature("regular")
	default:
		return setResponseLength("medium"), setTemperature("regular")
	}
}

func setTemperature(temperature string) float32 {
	switch temperature {
	case "creative":
		return 0.8
	case "deterministic":
		return 0.2
	case "regular":
		return 0.5
	default:
		return 0.5
	}
}

func setResponseLength(outputLength string) int32 {
	switch outputLength {
	case "short":
		return 250
	case "medium":
		return 500
	case "long":
		return 1200
	default:
		return 1000
	}
}

const MaxConversationLength = 10000 // Maximum length before summarization.

// calculateContextLength calculates the total length of the conversation.
func calculateContextLength(dialogues []entity.Dialogue) int {
	totalLength := 0
	for _, d := range dialogues {
		totalLength += len(d.Content)
	}
	return totalLength
}
