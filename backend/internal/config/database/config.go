package config

import (
	"backend/internal/config/encryption"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SQLConfig struct {
	Databases   map[string]*sql.DB
	SecurityKey []byte // Key for security-based encryption
	PIIKey      []byte // Key for PII encryption
}

func NewSQLConfig(dbConfigs map[string]string, securityKey, piiKey []byte) (*SQLConfig, error) {
	databases := make(map[string]*sql.DB)

	for dbName, dsn := range dbConfigs {
		db, err := connectAndCreateDatabase(dsn)
		if err != nil {
			return nil, err
		}

		log.Printf("Connected to database %s", dbName)
		databases[dbName] = db
	}

	return &SQLConfig{
		Databases:   databases,
		SecurityKey: securityKey,
		PIIKey:      piiKey,
	}, nil
}

func connectAndCreateDatabase(dsn string) (*sql.DB, error) {
	parts := strings.Split(dsn, "/")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid DSN format")
	}
	dbName := strings.Split(parts[1], "?")[0]

	dsnWithoutDB := strings.Join(parts[:len(parts)-1], "/") + "/"
	rootDB, err := sql.Open("mysql", dsnWithoutDB)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MySQL: %w", err)
	}
	defer rootDB.Close()

	_, err = rootDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		return nil, fmt.Errorf("error creating database %s: %w", dbName, err)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database %s: %w", dbName, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verifying connection to database %s: %w", dbName, err)
	}

	return db, nil
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

	tbl, ok := Tables[table]
	if !ok {
		return 0, fmt.Errorf("table %s not found in Tables map", table)
	}

	for i, col := range columns {
		// Security Encryption: Encrypt sensitive data
		if shouldEncryptColumn(tbl.SecurityEncryption, col) {
			if strVal, ok := values[i].(string); ok {
				encryptedVal, err := encryption.EncryptAES256(strVal, config.SecurityKey)
				if err != nil {
					return 0, fmt.Errorf("error encrypting value for column '%s': %w", col, err)
				}
				values[i] = encryptedVal
			}
		}

		// PII Encryption: Encrypt personally identifiable information
		if shouldEncryptColumn(tbl.PIIEncryption, col) {
			if strVal, ok := values[i].(string); ok {
				encryptedVal, err := encryption.EncryptAES128(strVal, config.PIIKey)
				if err != nil {
					return 0, fmt.Errorf("error encrypting value for column '%s': %w", col, err)
				}
				values[i] = encryptedVal
			}
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
	return strings.Join(elements, sep)
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

// Helper function to determine if a column should be encrypted
func shouldEncryptColumn(encryptedCols []string, column string) bool {
	for _, col := range encryptedCols {
		if col == column {
			return true
		}
	}
	return false
}
