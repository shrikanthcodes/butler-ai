package secrets

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenAI_Key() string {
	envPath := filepath.Join("secrets", "secret.env")
	err := LoadEnvFile(envPath)
	if err != nil {
		fmt.Printf("Error loading .env file")
		fmt.Println("Error extracting API key")
		return ""
	}
	apiKey := os.Getenv("OPENAI_KEY")
	return apiKey
}
