package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConfig struct {
	Databases map[string]*sql.DB
	SecretKey []byte // Key for AES encryption
}

func NewSQLConfig(dbConfigs map[string]string, encryptedDBs map[string]bool, secretKey []byte) (*SQLConfig, error) {
	databases := make(map[string]*sql.DB)

	for dbName, dsn := range dbConfigs {
		db, err := connectAndCreateDatabase(dsn)
		if err != nil {
			return nil, err
		}

		if encryptedDBs[dbName] {
			log.Printf("Database %s is encrypted", dbName)
		}

		log.Printf("Connected to database %s", dbName)
		databases[dbName] = db
	}

	return &SQLConfig{Databases: databases, SecretKey: secretKey}, nil
}

func connectAndCreateDatabase(dsn string) (*sql.DB, error) {
	// Extract the database name from the DSN
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid DSN format")
	}
	dbName := strings.Split(parts[1], "?")[0]

	// Create a DSN without the database name
	dsnWithoutDB := strings.Join(parts[:len(parts)-1], "/") + "/"

	// Connect to MySQL without specifying a database
	rootDB, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %w", err)
	}
	defer rootDB.Close()

	// Create the database if it doesn't exist
	_, err = rootDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return nil, fmt.Errorf("error creating database %s: %w", dbName, err)
	}

	// Now connect to the specific database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %s: %w", dbName, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to database %s: %w", dbName, err)
	}

	return db, nil
}

// AES encryption
func (config *SQLConfig) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(config.SecretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AES decryption
func (config *SQLConfig) Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(config.SecretKey)
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

func (config *SQLConfig) CreateTable(dbName, tableName, schema string) error {
	db, exists := config.Databases[dbName]
	if !exists {
		return fmt.Errorf("database %s not found", dbName)
	}

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("error creating table '%s': %w", tableName, err)
	}

	log.Printf("Table '%s' created successfully in database %s", tableName, dbName)
	return nil
}

func (config *SQLConfig) Query(dbName, sqlQuery string, params ...interface{}) (*sql.Rows, error) {
	db, exists := config.Databases[dbName]
	if !exists {
		return nil, fmt.Errorf("database %s not found", dbName)
	}

	rows, err := db.Query(sqlQuery, params...)
	if err != nil {
		return nil, fmt.Errorf("error executing query '%s': %w", sqlQuery, err)
	}
	return rows, nil
}

func (config *SQLConfig) Insert(dbName, table string, columns []string, values ...interface{}) (int64, error) {
	db, exists := config.Databases[dbName]
	if !exists {
		return 0, fmt.Errorf("database %s not found", dbName)
	}

	// Encrypt values before inserting
	for i, val := range values {
		if strVal, ok := val.(string); ok {
			encryptedVal, err := config.Encrypt(strVal)
			if err != nil {
				return 0, fmt.Errorf("error encrypting value: %w", err)
			}
			values[i] = encryptedVal
		}
	}

	columnsStr := fmt.Sprintf("`%s`", join(columns, "`, `"))
	placeholders := placeholders(len(values))
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, columnsStr, placeholders)

	result, err := db.Exec(query, values...)
	if err != nil {
		return 0, fmt.Errorf("error inserting row into table '%s': %w", table, err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}

	log.Printf("Inserted row into table '%s': %v, last insert ID: %d", table, values, lastID)
	return lastID, nil
}

func join(elements []string, sep string) string {
	result := ""
	for i, element := range elements {
		if i > 0 {
			result += sep
		}
		result += element
	}
	return result
}

func placeholders(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 {
			result += ", "
		}
		result += "?"
	}
	return result
}
