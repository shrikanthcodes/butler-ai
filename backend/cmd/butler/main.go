package main

import (
	config "backend/internal/config/database"
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
	// Load environment variables from .env file
	envPath := filepath.Join("internal", "constants", "db_credentials.env")
	if err := godotenv.Load(envPath); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Fetch database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	securityKeyBase64 := os.Getenv("SECURITY_KEY")
	piiKeyBase64 := os.Getenv("PII_KEY")

	// Validate the required environment variables
	if dbUser == "" || dbPass == "" || dbHost == "" || dbPort == "" || securityKeyBase64 == "" || piiKeyBase64 == "" {
		return fmt.Errorf("missing required environment variables")
	}

	// Decode the base64-encoded security key (AES-256 for security encryption)
	securityKey, err := base64.StdEncoding.DecodeString(securityKeyBase64)
	if err != nil {
		return fmt.Errorf("invalid base64-encoded security key: %w", err)
	}

	// Validate security key length (32 bytes for AES-256)
	if len(securityKey) != 32 {
		return fmt.Errorf("invalid AES-256 key size: %d bytes", len(securityKey))
	}

	// Decode the base64-encoded PII key (AES-128 for PII encryption)
	piiKey, err := base64.StdEncoding.DecodeString(piiKeyBase64)
	if err != nil {
		return fmt.Errorf("invalid base64-encoded PII key: %w", err)
	}

	// Validate PII key length (16 bytes for AES-128)
	if len(piiKey) != 16 {
		return fmt.Errorf("invalid AES-128 key size: %d bytes", len(piiKey))
	}

	// Construct the base DSN string
	baseDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)

	// Define the database configurations
	dbConfigs := map[string]string{
		"auth_db":   baseDSN + "auth_db?parseTime=true",
		"user_db":   baseDSN + "user_db?parseTime=true",
		"butler_db": baseDSN + "butler_db?parseTime=true",
	}

	// Initialize the SQL configuration with security and PII keys
	sqlConfig, err := config.NewSQLConfig(dbConfigs, securityKey, piiKey)
	if err != nil {
		return fmt.Errorf("failed to initialize SQLConfig: %w", err)
	}

	// Load schemas and test databases
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
