package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/internal/service/database"
	"github.com/shrikanthcodes/butler-ai/internal/service/llm"
	"github.com/shrikanthcodes/butler-ai/internal/service/repository"
	"github.com/shrikanthcodes/butler-ai/internal/service/templates"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
)

//Save conversation from repository to db is based on is_active:
//4 ways this is possible
//1. User ends the chat (is_active set to False, chat loaded to rabbitMQ),
//2. Same conversation is opened in another session
//	(is_active momentarily set to False (chat loaded to rabbitMQ) and is_active is true async,
//	3. Timeout (is_active set to false after a given amount of inactive time),
//	4. Refresh state (every 5 minutes, chat is backed up regardless of is_active and chat loaded to rabbitMQ)
//Implement Redis for in-memory repository, and RabbitMQ for message queueing

// CsService handles interactions with the chatbot.
type CsService struct {
	GeminiService     *llm.GsService
	DatabaseService   *database.DbService
	RepositoryService *repository.RsService
	TemplateService   *templates.TsService
	log               *logger.Logger
}

// NewChatService initializes the CsService with a repository
func NewChatService(geminiService *llm.GsService, templateService *templates.TsService,
	databaseService *database.DbService, repositoryService *repository.RsService,
	log *logger.Logger) (*CsService, error) {

	return &CsService{
		GeminiService:     geminiService,
		DatabaseService:   databaseService,
		RepositoryService: repositoryService,
		TemplateService:   templateService,
		log:               log,
	}, nil
}

// GetConversation retrieves 6 most recent dialogues of the conversation from transcript.
func getRecentDialogues() []entity.Dialogue {
	// Get 6 recent dialogues from the active conversation just loaded
	// from the db
	return []entity.Dialogue{}
}

// EndChat ends the conversation session and saves it to the db.
func (cs *CsService) EndChat(convID string) error {
	conversation.IsActive = false

	// Save the conversation to the db.
	err = cs.SaveConversation(conversation)
	if err != nil {
		return fmt.Errorf("failed to save conversation: %w", err)
	}

	cs.GeminiService.EndChat()
	return nil
}

// PredictChat adds a user message, generates a response from the AI, and updates the conversation.
func (cs *CsService) PredictChat(ctx context.Context, convID, chatType, userMessage string) (string, error) {
	conversation, err := cs.getConversation(convID)
	if err != nil {
		return "", err
	}

	// Generate response using GsService.
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
func (cs *CsService) StartChat(convID string, chatType string) error {
	templateName, maxTokens, temperature, err := MapAIParametersToChatType(chatType)
	if err != nil {
		return err
	}

	// Render the prompt using the templates service, data is the interface with all the variables
	prompt, err := cs.TemplateService.RenderTemplate(templateName, cs.DataBuilder(chatType))
	if err != nil {
		return fmt.Errorf("failed to render templates: %w", err)
	}

	conversation := entity.Conversation{}

	if _, exists := cs.ChatCache.activeConversations[convID]; !exists {
		conversation, err := cs.getConversation(convID)
		if err != nil {
			return fmt.Errorf("failed to load conversation: %w", err)
		}
		cs.ChatCache.activeConversations[convID] = conversation
	}

	recentDialogue := getRecentDialogues()

	if conversation.IsActive {
		err := cs.SaveConversation(conversation)
		if err != nil {
			return fmt.Errorf("failed to save conversation: %w", err)
		}
		//send conversation, summary from repository to db
		//start a new chat (not really) and reuse the cached conversation from in memory
	}
	cs.GeminiService.SetSystemPrompt(prompt)
	cs.GeminiService.SetModelParameters(maxTokens, temperature)
	// Set Conversation to active

	err = cs.GeminiService.StartNewChat(recentDialogue)
	if err != nil {
		return fmt.Errorf("failed to start new chat: %w", err)
	}

	conversation.IsActive = true
	cs.ChatCache.activeConversations[convID].IsActive = true

	return nil
}

// Save upon repository invalidation, implement Redis for in-memory repository

// getConversation retrieves a conversation from the db.
func (cs *CsService) getConversation(convID string) (*entity.Conversation, error) {
	conversation, err := cs.conversationDB.GetConversation(convID)
	if err != nil {
		return nil, err
	}
	return &conversation, nil
}

// updateConversationHistory updates the conversation history.
func (cs *CsService) updateConversationHistory(conversation *entity.Conversation, userDialogue, modelDialogue entity.Dialogue) {
	conversation.Conversation = append(conversation.Conversation, userDialogue, modelDialogue)

	// Maintain recent dialogues for summarization purposes (last 6 dialogues).
	conversation.RecentDialogues = append(conversation.RecentDialogues, userDialogue, modelDialogue)
	if len(conversation.RecentDialogues) > 6 {
		conversation.RecentDialogues = conversation.RecentDialogues[len(conversation.RecentDialogues)-6:]
	}
}

func (cs *CsService) CloseGeminiService() {
	err := cs.GeminiService.Close()
	if err != nil {
		fmt.Println("Failed to close GsService")
	}

}

// getTemplateNameByChatType returns the templates name based on the chat type.
func getTemplateNameByChatType(chatType string) (string, error) {
	switch chatType {
	case "recipe":
		return templates.RecipeMode, nil
	case "shopping":
		return templates.ShoppingMode, nil
	case "health":
		return templates.HealthMode, nil
	case "motivation":
		return templates.MotivationMode, nil
	case "calorie_tracker":
		return templates.CalorieTrackerMode, nil
	case "summarization":
		return templates.ChatSummarizationMode, nil
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

const MaxConversationLength = 10000 // Maximum length before summarization.

// calculateContextLength calculates the total length of the conversation.
func calculateContextLength(dialogues []entity.Dialogue) int {
	totalLength := 0
	for _, d := range dialogues {
		totalLength += len(d.Content)
	}
	return totalLength
}
