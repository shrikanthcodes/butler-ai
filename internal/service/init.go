package service

import (
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/controller/api/v1"
	"github.com/shrikanthcodes/butler-ai/internal/service/chat"
	"github.com/shrikanthcodes/butler-ai/internal/service/database"
	"github.com/shrikanthcodes/butler-ai/internal/service/llm"
	"github.com/shrikanthcodes/butler-ai/internal/service/template"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
)

func Start(cfg *config.Config, log *logger.Logger) error {
	templateService, err := template.NewTemplateService(log)
	checkError("TemplateService", err, log)

	geminiService, err := llm.NewGeminiService(log)
	checkError("GeminiService", err, log)

	databaseService, err := database.NewDatabaseService(cfg.Postgres.URL, log)
	checkError("DatabaseService", err, log)
	defer databaseService.Close()

	_, err = chat.NewChatService(geminiService, templateService, databaseService, log)
	checkError("ChatService", err, log)

	// Register API routes
	_, err = v1.RegisterRoutes(cfg.HTTP.Port, cfg.CORS, log)
	checkError("API Server", err, log)

	if err != nil {
		return err
	}
	return nil
}

// CheckError is a template function for error logging
func checkError(component string, err error, log *logger.Logger) {
	if err != nil {
		log.Fatal("Failed to initialize %s: %v", component, err)
	}
}
