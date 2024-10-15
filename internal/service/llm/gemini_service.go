package llm

import (
	"context"
	"errors"
	"fmt"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"github.com/shrikanthcodes/butler-ai/internal/entity"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// GeminiService handles interactions with Gemini AI.
type GeminiService struct {
	client      *genai.Client
	model       *genai.GenerativeModel
	chatSession *genai.ChatSession
	log         *logger.Logger
	mu          sync.Mutex // Mutex to protect shared resources.
}

// NewGeminiService creates a new instance of GeminiService with safety settings.
func NewGeminiService(log *logger.Logger) (*GeminiService, error) {

	const ModelName = "gemini-1.5-flash"

	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("api key not found")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	model := client.GenerativeModel(ModelName)
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockLowAndAbove,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockLowAndAbove,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockLowAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	chatSession := model.StartChat()

	return &GeminiService{
		client:      client,
		model:       model,
		chatSession: chatSession,
		log:         log,
	}, nil
}

// Close gracefully shuts down the GeminiService.
func (gs *GeminiService) Close() error {
	return gs.client.Close()
}

// SetModelParameters sets the model parameters for the chat session.
func (gs *GeminiService) SetModelParameters(maxTokens int32, temperature float32) *GeminiService {
	gs.model.SetMaxOutputTokens(maxTokens)
	gs.model.SetTemperature(temperature)
	return gs
}

// SetSystemPrompt sets the system prompt for the chat session.
func (gs *GeminiService) SetSystemPrompt(systemPrompt string) *GeminiService {
	gs.model.SystemInstruction = genai.NewUserContent(genai.Text(systemPrompt))
	return gs

}

// StartNewChat starts a new chat session without previous context.
func (gs *GeminiService) StartNewChat(recentDialogues []entity.Dialogue) error {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	gs.chatSession = gs.model.StartChat()

	// Load recent dialogues into the chat history.
	for _, dialogue := range recentDialogues {
		gs.appendDialogueToChatHistory(dialogue.Role, dialogue.Content)
	}

	return nil
}

// EndChat ends the current chat session.
func (gs *GeminiService) EndChat() {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	log.Println("Closing Chat Session")
	gs.chatSession = nil
}

// appendDialogueToChatHistory adds a new dialogue to the chat session.
func (gs *GeminiService) appendDialogueToChatHistory(role, content string) {
	newContent := &genai.Content{
		Parts: []genai.Part{
			genai.Text(content),
		},
		Role: role,
	}
	gs.chatSession.History = append(gs.chatSession.History, newContent)
}

// PredictChat generates the next dialogue in a conversation.
func (gs *GeminiService) PredictChat(ctx context.Context, userMessage string) (string, error) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.chatSession == nil {
		return "", errors.New("can't start chat session without a system prompt")
	}

	// Add user message to history.
	gs.appendDialogueToChatHistory(entity.RoleUser, userMessage)

	// Send message and get response.
	response, err := gs.chatSession.SendMessage(ctx, genai.Text(userMessage))
	if err != nil {
		// Check if the error is a googleapi.Error and extract details.
		var googleErr *googleapi.Error
		if errors.As(err, &googleErr) {
			log.Printf("Google API error: Code %d, Message: %s, Details: %v\n", googleErr.Code, googleErr.Message, googleErr.Body)
		} else {
			log.Printf("Unexpected error: %v\n", err)
		}
		return "", fmt.Errorf("prediction error: %w", err)
	}

	// Extract text from response.
	var result strings.Builder
	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			if textPart, ok := part.(genai.Text); ok {
				result.WriteString(string(textPart))
			}
		}
	}

	// Append assistant's response to chat history.
	gs.appendDialogueToChatHistory(entity.RoleModel, result.String())

	return result.String(), nil
}

// Predict generates a one-shot response based on the provided text.
func (gs *GeminiService) Predict(ctx context.Context, text string, maxTokens int32, temperature float32) (string, error) {
	gs.mu.Lock()
	gs.model.SetMaxOutputTokens(maxTokens)
	gs.model.SetTemperature(temperature)
	gs.mu.Unlock()

	// Generate content.
	iter := gs.model.GenerateContentStream(ctx, genai.Text(text))
	var result strings.Builder
	for {
		response, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("generation error: %w", err)
		}
		// Extract text from response.
		for _, candidate := range response.Candidates {
			for _, part := range candidate.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					result.WriteString(string(textPart))
				}
			}
		}
	}
	return result.String(), nil
}
