package butlerai

import (
	"bufio"
	"fmt"
	config "github.com/shrikanthcodes/butler-ai/backend/config"
	chat "github.com/shrikanthcodes/butler-ai/backend/internal/services"
	templates "github.com/shrikanthcodes/butler-ai/backend/pkg/services/templates"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

func Run(cfg *config.Config) {
	xy, err := config.InitDB(cfg)
	if err != nil {
		log.Default()
	}

	fmt.Print(xy)

	templatesDir := filepath.Join("pkg", "services", "templates", "resources")
	templateService, err := templates.NewTemplateService(templatesDir)
	if err != nil {
		log.Fatalf("Failed to initialize TemplateService: %v", err)
	}

	var db *gorm.DB
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
