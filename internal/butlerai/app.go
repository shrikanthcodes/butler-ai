package butlerai

import (
	config "github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/service"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/postgres"
)

func Run(cfg *config.Config) {
	// Creates objects via constructors
	log := logger.New(cfg.Log.Level)
	log.Info("Starting Butler AI")

	pg, err := postgres.New(cfg.Postgres.URL, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		log.Fatal("Failed to initialize Postgres", err)
	}
	defer pg.Close()

	// Initialize TemplateService
	templateService, err := service.NewTemplateService()
	if err != nil {
		log.Fatal("Failed to initialize TemplateService", err)
	}

	// Initialize GeminiService
	geminiService, err := service.NewGeminiService()
	if err != nil {
		log.Fatal("Failed to initialize GeminiService", err)
	}

	// Initialize DBConversationStore
	dbStore := service.NewDBConversationStore(pg.DB)

	// Initialize ChatService with injected dependencies
	chatService := service.NewChatService(geminiService, templateService, dbStore)

	// Initialize API Server
	apiServer := api.NewServer(chatService, log)

	// Start the server
	if err := apiServer.Start(cfg.HTTP.Port); err != nil {
		log.Fatal("Failed to start API server", err)
	}
	//Initialize chat service which initializes gemini

	// Initialize services
	// InitializeChatService(db *DBConversationStore, templateService *TemplateService) (*ChatService, error)

	// InitializeChatService initializes the ChatService with the database store
	// func InitializeChatService(db *DBConversationStore, templateService *TemplateService) (*ChatService, error) {

	//templatesDir := filepath.Join("pkg", "services", "templates", "resources")
	//templateService, err := service.NewTemplateService(templatesDir)
	//if err != nil {
	//	log.Fatalf("Failed to initialize TemplateService: %v", err)
	//}
	//
	//var db *gorm.DB
	//conversationDB := service.NewDBConversationStore(db)
	//chatService, err := service.InitializeChatService(conversationDB, templateService)
	//if err != nil {
	//	log.Fatalf("Failed to initialize ChatService: %v", err)
	//}
	//
	//convID := "conversation123"
	//chatType := "recipe"
	//
	//err = chatService.StartChat(convID, chatType)
	//if err != nil {
	//	log.Fatalf("Failed to start chat: %v", err)
	//}
	//
	//fmt.Printf("Chat session started with ID '%s' and type '%s'.\n", convID, chatType)
	//
	//ctx := context.Background()
	//reader := bufio.NewReader(os.Stdin)
	//text := ""
	//
	//for text != "exit" {
	//	fmt.Println("Enter text: ")
	//	text, _ = reader.ReadString('\n')
	//
	//	aiResponse, err := chatService.PredictChat(ctx, convID, chatType, text+"(Answer in 50 words or less)")
	//	if err != nil {
	//		log.Printf("Failed to get AI response: %v", err)
	//	} else {
	//		fmt.Printf("Assistant: %s\n", aiResponse)
	//	}
	//}
	//
	//err = chatService.EndChat(convID)
	//if err != nil {
	//	log.Printf("Failed to end chat: %v", err)
	//}
	//fmt.Println("Chat session ended.")
	//
	//chatService.CloseGeminiService()
}
