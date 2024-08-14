package tests

import (
	"backend/internal/config/database"
	"fmt"
	"log"
	"testing"
)

// TestDBOperations tests the CRUD operations on the database.
func TestDBOperations(t *testing.T) {
	// Initialize your database configurations
	dbConfigs := map[string]string{
		"auth_db":   "your_dsn_for_auth_db",
		"user_db":   "your_dsn_for_user_db",
		"butler_db": "your_dsn_for_butler_db",
	}

	// Initialize your encryption keys
	securityKey := []byte("your_security_key")
	piiKey := []byte("your_pii_key")

	// Initialize the SQLConfig
	sqlConfig, err := database.NewSQLConfig(dbConfigs, securityKey, piiKey)
	if err != nil {
		t.Fatalf("Failed to initialize SQLConfig: %v", err)
	}

	// Load the schemas
	configLoader := database.NewConfigLoader(sqlConfig)
	if err := configLoader.LoadSchemas(); err != nil {
		t.Fatalf("Failed to load schemas: %v", err)
	}

	// Test each table by performing insert and query operations
	for _, tableName := range database.TableOrder {
		t.Run(fmt.Sprintf("TestTable_%s", tableName), func(t *testing.T) {
			if err := testTableOperations(sqlConfig, tableName); err != nil {
				t.Fatalf("Test failed for table %s: %v", tableName, err)
			}
		})
	}
}

// testTableOperations performs insert and query operations on a given table
func testTableOperations(sqlConfig *database.SQLConfig, tableName string) error {
	table := database.Tables[tableName]
	dbName := table.DbName

	// Prepare dummy data for insert operation
	values := make([]interface{}, len(table.Columns))
	for i := range values {
		switch table.Columns[i] {
		case "user_id", "refresh_token", "recipe_id":
			values[i] = []byte(fmt.Sprintf("test_%s_%d", tableName, i)) // Assuming VARBINARY fields
		default:
			values[i] = fmt.Sprintf("test_%s_%d", tableName, i)
		}
	}

	// Perform the insert operation
	_, err := sqlConfig.Insert(dbName, tableName, table.Columns, values...)
	if err != nil {
		return fmt.Errorf("error inserting into table %s: %w", tableName, err)
	}

	// Perform the query operation
	query := fmt.Sprintf("SELECT * FROM %s LIMIT 1", tableName)
	rows, err := sqlConfig.Query(dbName, query)
	if err != nil {
		return fmt.Errorf("error querying table %s: %w", tableName, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("no rows returned from table %s", tableName)
	}

	log.Printf("Table %s tested successfully", tableName)
	return nil
}
