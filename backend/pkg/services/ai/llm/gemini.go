package ai

import (
	context "context"
	errors "errors"
	fmt "fmt"
	log "log"
	os "os"
	strings "strings"

	models "backend/pkg/models"

	genai "github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi"
	iterator "google.golang.org/api/iterator"
	option "google.golang.org/api/option"
)

// GeminiService handles interactions with Gemini AI
type GeminiService struct {
	client         *genai.Client
	model          *genai.GenerativeModel
	chatSession    *genai.ChatSession
	safetySettings []*genai.SafetySetting
}

func get_api_key() string {
	err := godotenv.Load("/home/shrikanth/Documents/GitHub/butlerAI/backend/ai_secrets.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	API_KEY := os.Getenv("GEMINI_API_KEY")
	if API_KEY == "" {
		log.Fatalf("API_KEY not set in environment")
	}
	return API_KEY
}

const MODEL_NAME = "gemini-1.5-flash"

// InitializeGeminiService creates a new instance of GeminiService with safety settings
func InitializeGeminiService() *GeminiService {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(get_api_key()))
	if err != nil {
		log.Fatalf("Failed to create genai client: %v", err)
	}
	model := client.GenerativeModel(MODEL_NAME)
	// Configure default safety settings
	safetySettings := DefaultSafetySettings()

	chatSession := model.StartChat()

	return &GeminiService{
		client:         client,
		model:          model,
		chatSession:    chatSession,
		safetySettings: safetySettings,
	}
}

// Close gracefully shuts down the GeminiService
func (gs *GeminiService) Close() error {
	log.Println("Closing GeminiService")
	return gs.client.Close()
}

// StartNewChat starts a new chat session without previous context
func (gs *GeminiService) StartNewChat(prompt string, recent_dialogues []models.Dialogue, maxTokens int32, temperature float32) {
	if len(gs.chatSession.History) != 0 {
		log.Panic("There is an ongoing chat session")
	}
	gs.chatSession = gs.model.StartChat()

	gs.model.SetMaxOutputTokens(maxTokens)
	gs.model.SetTemperature(temperature)

	gs.model.SystemInstruction = genai.NewUserContent(genai.Text(prompt))

	// Load n recent dialogue from conversation history
	for _, dialogue := range recent_dialogues {
		gs.AppendDialogueToChatHistory(dialogue.Role, dialogue.Content)
	}
}

// EndChat ends the current chat session
func (gs *GeminiService) EndChat() {
	log.Println("Closing Chat Session")
	gs.chatSession = nil
}

// AppendDialogueToChatHistory adds a new dialogue to the chat session
func (gs *GeminiService) AppendDialogueToChatHistory(role, content string) {
	newContent := &genai.Content{
		Parts: []genai.Part{
			genai.Text(content),
		},
		Role: role,
	}
	gs.chatSession.History = append(gs.chatSession.History, newContent)
}

// PredictChat generates the next dialogue in a conversation
func (gs *GeminiService) PredictChat(ctx context.Context, userMessage string) (string, error) {
	// Start a new chat session with a new context
	if gs.chatSession == nil {
		log.Panic("Can't start chat session without a system prompt")
	}

	// Add user message to history
	gs.AppendDialogueToChatHistory(models.SetRole("user"), userMessage)

	// Send message and get response
	response, err := gs.chatSession.SendMessage(ctx, genai.Text(userMessage))
	if err != nil {
		// Check if the error is a googleapi.Error and extract details
		var googleErr *googleapi.Error
		if errors.As(err, &googleErr) {
			fmt.Printf("Google API error: Code %d, Message: %s, Details: %v\n", googleErr.Code, googleErr.Message, googleErr.Body)
		} else {
			fmt.Printf("Unexpected error: %v\n", err)
		}
		return "", fmt.Errorf("prediction error: %w", err)
	}

	// Extract text from response
	var result strings.Builder
	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			if textPart, ok := part.(genai.Text); ok {
				result.WriteString(string(textPart))
			}
		}
	}

	// Append assistant's response to chat history
	gs.AppendDialogueToChatHistory(models.SetRole("model"), result.String())

	return result.String(), nil
}

// Predict generates a one-shot response based on the provided text
func (gs *GeminiService) Predict(ctx context.Context, text string, maxTokens int32, temperature float32) (string, error) {
	// Set model parameters
	gs.model.SetMaxOutputTokens(maxTokens)
	gs.model.SetTemperature(temperature)

	// Generate content
	iter := gs.model.GenerateContentStream(ctx, genai.Text(text))
	var result strings.Builder
	for {
		response, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Extract text from response
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

// DefaultSafetySettings allows setting the safety thresholds
func DefaultSafetySettings() []*genai.SafetySetting {
	// Configure default safety settings
	safetySettings := []*genai.SafetySetting{
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
	return safetySettings
}
