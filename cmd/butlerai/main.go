package main

import (
	"log"

	config "github.com/shrikanthcodes/butler-ai/config"
	butlerai "github.com/shrikanthcodes/butler-ai/internal/butlerai"
)

func main() {
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Run
	butlerai.Run(cfg)
}
