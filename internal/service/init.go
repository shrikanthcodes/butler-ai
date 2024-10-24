package service

import (
	"github.com/shrikanthcodes/butler-ai/config"
	v1 "github.com/shrikanthcodes/butler-ai/internal/controller/api/v1"
	"github.com/shrikanthcodes/butler-ai/internal/service/cache"
	"github.com/shrikanthcodes/butler-ai/internal/service/chat"
	"github.com/shrikanthcodes/butler-ai/internal/service/database"
	"github.com/shrikanthcodes/butler-ai/internal/service/llm"
	"github.com/shrikanthcodes/butler-ai/internal/service/queue"
	"github.com/shrikanthcodes/butler-ai/internal/service/templates"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
)

func Start(cfg *config.Config, log *logger.Logger) error {
	// Register API routes
	_, err := v1.RegisterRoutes(cfg.HTTP.Port, cfg.CORS, log)
	checkError("APIServer", err, log)
	defer deferClose("APIServer", v1.Close, log)

	// Initialize TemplateService
	templateService, err := templates.NewTemplateService(log)
	checkError("TemplateService", err, log)
	defer deferClose("TemplateService", templateService.Close, log)

	// Initialize GeminiService
	geminiService, err := llm.NewGeminiService(log)
	checkError("GeminiService", err, log)
	defer deferClose("GeminiService", geminiService.Close, log)

	// Initialize DatabaseService
	databaseService, err := database.NewDatabaseService(cfg.Postgres.URL, log)
	checkError("DatabaseService", err, log)
	defer deferClose("DatabaseService", databaseService.Close, log)

	// Initialize CacheService
	cacheService, err := cache.NewCacheService(cfg.Redis, log)
	checkError("CacheService", err, log)
	defer deferClose("RepositoryService", cacheService.Close, log)

	// Initialize RepositoryService
	queueService, err := queue.NewQueueService(cfg.RabbitMQ, log)
	checkError("QueueService", err, log)
	defer deferClose("QueueService", queueService.Close, log)

	// Initialize ChatService
	csBuilder := &chat.CsService{
		GeminiService:   geminiService,
		DatabaseService: databaseService,
		CacheService:    cacheService,
		TemplateService: templateService,
		QueueService:    queueService,
		Log:             log,
	}
	_, err = chat.NewChatService(csBuilder)
	checkError("ChatService", err, log)

	if err != nil {
		return err
	}
	return nil
}

// checkError is a helper function for error logging.
func checkError(component string, err error, log *logger.Logger) {
	if err != nil {
		log.Fatal("Failed to initialize %s: %v", component, err)
	}
}

// deferClose is a helper function for handling deferred closures.
func deferClose(component string, closeFunc func() error, log *logger.Logger) {
	err := closeFunc()
	if err != nil {
		log.Error("Failed to close %s: %v", component, err)
	}
}
