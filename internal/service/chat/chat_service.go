package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"github.com/shrikanthcodes/butler-ai/internal/service/database"
	"github.com/shrikanthcodes/butler-ai/internal/service/llm"
	"github.com/shrikanthcodes/butler-ai/internal/service/template"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"sync"
	"time"
)

//Save conversation from cache to db is based on is_active:
//4 ways this is possible
//1. User ends the chat (is_active set to False, chat loaded to rabbitMQ),
//2. Same conversation is opened in another session
//	(is_active momentarily set to False (chat loaded to rabbitMQ) and is_active is true async,
//	3. Timeout (is_active set to false after a given amount of inactive time),
//	4. Refresh state (every 5 minutes, chat is backed up regardless of is_active and chat loaded to rabbitMQ)
//Implement Redis for in-memory cache, and RabbitMQ for message queueing

// ChatService handles interactions with the chatbot.
type ChatService struct {
	GeminiService   *llm.GeminiService
	DatabaseService *database.DatabaseService
	ChatCache       *ChatCacheRepository
	TemplateService *template.TemplateService
	log             *logger.Logger
}

// ChatCacheRepository stores active conversations in-memory for batching
type ChatCacheRepository struct {
	activeConversations map[string]*entity.Conversation
	recentDialogues     map[string]*[]entity.Dialogue
	lastUpdated         map[string]time.Time
	mu                  sync.Mutex
}

func NewChatCache() (*ChatCacheRepository, error) {
	return &ChatCacheRepository{
		activeConversations: make(map[string]*entity.Conversation),
		recentDialogues:     make(map[string]*[]entity.Dialogue),
		lastUpdated:         make(map[string]time.Time),
	}, nil
}

// NewChatService initializes the ChatService with a cache
func NewChatService(geminiService *llm.GeminiService, templateService *template.TemplateService,
	databaseService *database.DatabaseService, log *logger.Logger) (*ChatService, error) {
	cache, err := NewChatCache()
	if err != nil {
		return nil, err
	}
	return &ChatService{
		GeminiService:   geminiService,
		DatabaseService: databaseService,
		ChatCache:       cache,
		TemplateService: templateService,
		log:             log,
	}, nil
}

// GetConversation retrieves 6 most recent dialogues of the conversation from transcript.
func getRecentDialogues() []entity.Dialogue {
	// Get 6 recent dialogues from the active conversation just loaded
	// from the db
	return []entity.Dialogue{}
}

// EndChat ends the conversation session and saves it to the db.
func (cs *ChatService) EndChat(convID string) error {
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
	templateName, maxTokens, temperature, err := MapAIParametersToChatType(chatType)
	if err != nil {
		return err
	}

	// Render the prompt using the template service, data is the interface with all the variables
	prompt, err := cs.TemplateService.RenderTemplate(templateName, cs.DataBuilder(chatType))
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
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
		//send conversation, summary from cache to db
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

// Save upon cache invalidation, implement Redis for in-memory cache

// getConversation retrieves a conversation from the db.
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
	err := cs.GeminiService.Close()
	if err != nil {
		fmt.Println("Failed to close GeminiService")
	}

}

// getTemplateNameByChatType returns the template name based on the chat type.
func getTemplateNameByChatType(chatType string) (string, error) {
	switch chatType {
	case "recipe":
		return template.RecipeMode, nil
	case "shopping":
		return template.ShoppingMode, nil
	case "health":
		return template.HealthMode, nil
	case "motivation":
		return template.MotivationMode, nil
	case "calorie_tracker":
		return template.CalorieTrackerMode, nil
	case "summarization":
		return template.ChatSummarizationMode, nil
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
