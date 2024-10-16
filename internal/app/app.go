package app

import (
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/service"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
)

// Run is responsible for initializing objects
func Run(cfg *config.Config) {
	// Initialize logger
	log := logger.New(cfg.Log.Level)
	log.Info("Starting Butler AI")

	// Initialize services
	err := service.Start(cfg, log)
	if err != nil {
		log.Fatal("Failed to initialize services", err)
	}
	log.Info("Butler AI started successfully, listening on port: ", cfg.HTTP.Port)
}
