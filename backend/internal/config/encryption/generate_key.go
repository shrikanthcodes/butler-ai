package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	secretKey, err := generateSecretKey(32) // 32 bytes = 256 bits
	if err != nil {
		log.Fatalf("Error generating secret key: %v", err)
	}

	fmt.Printf("Generated Secret Key: %s\n", secretKey)
}

func generateSecretKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
