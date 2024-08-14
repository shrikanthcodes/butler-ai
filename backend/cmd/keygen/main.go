package main

import (
	"backend/internal/config/encryption"
)

func main() {
	// Generate AES-256 key
	encryption.GenerateAES256Key()

	// Generate AES-128 key
	encryption.GenerateAES128Key()
}
