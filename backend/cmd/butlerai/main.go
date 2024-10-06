package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"backend/pkg/models"
	"backend/pkg/services/chat"
	"backend/pkg/services/templates"

	"github.com/joho/godotenv"
)

// InMemoryConversationDB is a simple in-memory implementation of ConversationDB for testing.
type InMemoryConversationDB struct {
	conversations map[string]*chat.Conversation
	mu            sync.Mutex
}

// NewInMemoryConversationDB initializes the InMemoryConversationDB.
func NewInMemoryConversationDB() *InMemoryConversationDB {
	return &InMemoryConversationDB{
		conversations: make(map[string]*chat.Conversation),
	}
}

// GetConversation retrieves an existing conversation by ID or creates a new one if it doesn't exist.
func (db *InMemoryConversationDB) GetConversation(convID string) (*chat.Conversation, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Check if the conversation already exists in memory
	if conv, ok := db.conversations[convID]; ok {
		return conv, nil
	}

	// If the conversation does not exist, create a new one
	conv := &chat.Conversation{
		ID:               convID,
		FullConversation: []models.Dialogue{},
		RecentDialogue:   []models.Dialogue{},
		Summary:          "",
		IsActive:         false,
	}

	// Save the new conversation in the in-memory store
	db.conversations[convID] = conv
	return conv, nil
}

// SaveConversation stores or updates a conversation in the in-memory database.
func (db *InMemoryConversationDB) SaveConversation(conv *chat.Conversation) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.conversations[conv.ID] = conv
	return nil
}

// DataBuilder builds the data context for templates based on the chatType.
func (db *InMemoryConversationDB) DataBuilder(chatType string, outputLength int32) map[string]interface{} {
	// For testing purposes, return some fake data depending on the chat type.
	switch chatType {
	case "recipe":
		return map[string]interface{}{
			"UserName":     "TestUser",
			"OutputLength": outputLength,
			"Dish":         "Spaghetti Bolognese",
			"Ingredients":  []string{"spaghetti", "minced meat", "tomato sauce", "onion", "garlic"},
		}
	case "shopping":
		return map[string]interface{}{
			"UserName":     "TestUser",
			"OutputLength": outputLength,
			"ShoppingList": []string{"apples", "bananas", "milk", "bread"},
		}
	case "health":
		return map[string]interface{}{
			"UserName":           "TestUser",
			"OutputLength":       outputLength,
			"FitnessGoal":        "Lose weight",
			"PreferredExercises": []string{"running", "yoga", "cycling"},
		}
	case "motivation":
		return map[string]interface{}{
			"UserName":          "TestUser",
			"OutputLength":      outputLength,
			"MotivationMessage": "You are capable of achieving anything you set your mind to.",
		}
	case "calorie_tracker":
		return map[string]interface{}{
			"UserName":         "TestUser",
			"OutputLength":     outputLength,
			"CaloriesConsumed": 1200,
			"TargetCalories":   2000,
		}
	default:
		return nil
	}
}

func main() {
	fmt.Println("Loading .env file...")
	// Load environment variables from .env file
	err := godotenv.Load("ai_secrets.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println(".env file loaded successfully.")
	fmt.Println("Initializing TemplateService...")
	// Initialize TemplateService with the path to your templates directory.
	templatesDir := filepath.Join("pkg", "services", "templates", "resources")
	templateService, err := templates.NewTemplateService(templatesDir)
	if err != nil {
		log.Fatalf("Failed to initialize TemplateService: %v", err)
	}

	fmt.Println("Initializing InMemoryConversationDB...")

	// Initialize InMemoryConversationDB for testing.
	conversationDB := NewInMemoryConversationDB()

	fmt.Println("Initializing ChatService...")

	// Initialize ChatService with dependencies.
	chatService, err := chat.InitializeChatService(conversationDB, templateService)
	if err != nil {
		log.Fatalf("Failed to initialize ChatService: %v", err)
	}

	fmt.Println("ChatService initialized successfully.")

	// Use a unique conversation ID for testing.
	convID := "test-conversation-1"

	// Start a chat session with a specific chat type.
	chatType := "recipe" // Can be "recipe", "shopping", "health", etc.
	err = chatService.StartChat(convID, chatType)
	if err != nil {
		log.Fatalf("Failed to start chat: %v", err)
	}

	fmt.Printf("Chat session started with ID '%s' and type '%s'.\n", convID, chatType)

	// Create a buffered reader for user input.
	reader := bufio.NewReader(os.Stdin)

	// Context for API calls.
	ctx := context.Background()

	// Simulate a conversation loop.
	for {
		fmt.Print("You: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read user input: %v", err)
		}

		// Trim the newline character from user input.
		userInput = userInput[:len(userInput)-1]

		if userInput == "exit" {
			// End the chat session.
			err = chatService.EndChat(convID)
			if err != nil {
				log.Printf("Failed to end chat: %v", err)
			}
			fmt.Println("Chat session ended.")
			break
		}

		// Get AI response.
		aiResponse, err := chatService.PredictChat(ctx, convID, chatType, userInput)
		if err != nil {
			log.Printf("Failed to get AI response: %v", err)
			continue
		}

		fmt.Printf("Assistant: %s\n", aiResponse)
	}

	// Close the GeminiService when done.
	err = chatService.GeminiService().Close()
	if err != nil {
		log.Printf("Failed to close GeminiService: %v", err)
	}
}
