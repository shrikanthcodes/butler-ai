package main

import (
	"log"

	butlerai "github.com/shrikanthcodes/butler-ai/backend/internal/butlerai"
	config "github.com/shrikanthcodes/butler-ai/backend/internal/config"
)

func main() {
	cfg = "string"
	cfg, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Run
	butlerai.Run(cfg)
}
