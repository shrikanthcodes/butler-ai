package encryption

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// GenerateAES256Key generates a 32-byte AES-256 key.
func GenerateAES256Key() {
	secretKey, err := generateSecretKey(32) // 32 bytes = 256 bits
	if err != nil {
		log.Fatalf("Error generating AES-256 secret key: %v", err)
	}
	fmt.Printf("Generated AES-256 Secret Key: %s\n", secretKey)
}

// GenerateAES128Key generates a 16-byte AES-128 key.
func GenerateAES128Key() {
	secretKey, err := generateSecretKey(16) // 16 bytes = 128 bits
	if err != nil {
		log.Fatalf("Error generating AES-128 secret key: %v", err)
	}
	fmt.Printf("Generated AES-128 Secret Key: %s\n", secretKey)
}

// generateSecretKey generates a base64-encoded secret key of the given length.
func generateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
