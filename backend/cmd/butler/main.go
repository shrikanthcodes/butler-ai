package main

import (
	"backend/internal/config"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	envPath := filepath.Join("internal", "constants", "db_credentials.env")
	if err := godotenv.Load(envPath); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	secretKeyBase64 := os.Getenv("SECRET_KEY")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || secretKeyBase64 == "" {
		return fmt.Errorf("missing required environment variables")
	}

	// Decode the base64-encoded secret key
	secretKey, err := base64.StdEncoding.DecodeString(secretKeyBase64)
	if err != nil {
		return fmt.Errorf("invalid base64-encoded secret key: %w", err)
	}

	// Validate the key length
	if len(secretKey) != 16 && len(secretKey) != 24 && len(secretKey) != 32 {
		return fmt.Errorf("invalid AES key size: %d bytes", len(secretKey))
	}

	baseDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)

	dbConfigs := map[string]string{
		"auth_db":   baseDSN + "auth_db?parseTime=true",
		"user_db":   baseDSN + "user_db?parseTime=true",
		"butler_db": baseDSN + "butler_db?parseTime=true",
	}

	encryptedDBs := map[string]bool{
		"auth_db":   true,
		"user_db":   false,
		"butler_db": false,
	}

	sqlConfig, err := config.NewSQLConfig(dbConfigs, encryptedDBs, secretKey)
	if err != nil {
		return fmt.Errorf("failed to initialize SQLConfig: %w", err)
	}

	configLoader := config.NewConfigLoader(sqlConfig)
	if err := configLoader.LoadSchemas(); err != nil {
		return fmt.Errorf("failed to load schemas: %w", err)
	}

	if err := configLoader.TestDatabases(); err != nil {
		return fmt.Errorf("database tests failed: %w", err)
	}

	log.Println("All database schemas loaded and tested successfully")
	return nil
}
