package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// DecryptAES256 decrypts a ciphertext string using AES-256 encryption.
func DecryptAES256(ciphertext string, secretKey []byte) (string, error) {
	if len(secretKey) != 32 {
		return "", errors.New("invalid key size for AES-256")
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertextBytes := decodedData[:nonceSize], decodedData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// DecryptAES128 decrypts a ciphertext string using AES-128 encryption (simpler for PII).
func DecryptAES128(ciphertext string, secretKey []byte) (string, error) {
	if len(secretKey) != 16 {
		return "", errors.New("invalid key size for AES-128")
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decodedData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertextBytes := decodedData[:nonceSize], decodedData[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
