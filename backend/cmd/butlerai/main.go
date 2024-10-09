package main

import (
	"backend/internal/config"
	"backend/pkg/services/chat"
	"backend/pkg/services/templates"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// fmt.Println("Loading .env file...")
	// err := godotenv.Load("ai_secrets.env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }
	// Create a new conversation

	db, err := config.InitDB()
	if err != nil {
		log.Default()
	}

	// conv := &models.Conversation{
	// 	ConvID: "conversation123",
	// 	UserID: "user123",
	// 	Conversation: DialogueArray{
	// 		{Role: models.RoleUser, Content: "Winter recipe pls"},
	// 		{Role: models.RoleModel, Content: "Sure!"},
	// 	},
	// 	IsActive: true,
	// }
	// err = db.Save(&conv).Error
	// if err != nil {
	// 	fmt.Println("Error saving conversation:", err)
	// }

	templatesDir := filepath.Join("pkg", "services", "templates", "resources")
	templateService, err := templates.NewTemplateService(templatesDir)
	if err != nil {
		log.Fatalf("Failed to initialize TemplateService: %v", err)
	}

	conversationDB := chat.NewDBConversationStore(db)
	chatService, err := chat.InitializeChatService(conversationDB, templateService)
	if err != nil {
		log.Fatalf("Failed to initialize ChatService: %v", err)
	}

	convID := "conversation123"
	chatType := "recipe"

	err = chatService.StartChat(convID, chatType)
	if err != nil {
		log.Fatalf("Failed to start chat: %v", err)
	}

	fmt.Printf("Chat session started with ID '%s' and type '%s'.\n", convID, chatType)

	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)
	text := ""

	for text != "exit" {
		fmt.Println("Enter text: ")
		text, _ = reader.ReadString('\n')

		aiResponse, err := chatService.PredictChat(ctx, convID, chatType, text+"(Answer in 50 words or less)")
		if err != nil {
			log.Printf("Failed to get AI response: %v", err)
		} else {
			fmt.Printf("Assistant: %s\n", aiResponse)
		}
	}

	err = chatService.EndChat(convID)
	if err != nil {
		log.Printf("Failed to end chat: %v", err)
	}
	fmt.Println("Chat session ended.")

	chatService.CloseGeminiService()

}
