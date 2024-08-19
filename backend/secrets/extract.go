package secrets

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenAI_Key() string {
	envPath := filepath.Join("secrets", ".env")
	err := LoadEnvFile(envPath)
	if err != nil {
		fmt.Printf("Error loading .env file from location: %v", envPath)
		fmt.Println("Error extracting API key")
		return ""
	}
	apiKey := os.Getenv("OPENAI_KEY")
	return apiKey
}
