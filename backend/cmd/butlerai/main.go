package main

import (
	"log"

	config "github.com/shrikanthcodes/butler-ai/backend/config"
	butlerai "github.com/shrikanthcodes/butler-ai/backend/internal/butlerai"
)

func main() {
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Run
	butlerai.Run(cfg)
}
