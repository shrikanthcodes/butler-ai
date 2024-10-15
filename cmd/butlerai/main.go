package main

import (
	"log"

	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/internal/app"
)

func main() {
	cfg, err := config.SetConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Run
	app.Run(cfg)
}
