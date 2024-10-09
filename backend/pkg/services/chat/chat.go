package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/backend/internal/entity"

	ai "github.com/shrikanthcodes/butler-ai/backend/pkg/services/ai/llm"
	"github.com/shrikanthcodes/butler-ai/backend/pkg/services/templates"
)

type ChatService struct {
	geminiService   *ai.GeminiService
	templateService *templates.TemplateService
	conversationDB  *DBConversationStore
}

// InitializeChatService initializes the ChatService with the database store
func InitializeChatService(db *DBConversationStore, templateService *templates.TemplateService) (*ChatService, error) {
	gemini, err := ai.InitializeGeminiService()
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
