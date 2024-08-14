// configLoader.go

package config

import (
	"fmt"
	"log"
)

type ConfigLoader struct {
	SQLConfig *SQLConfig
}

func NewConfigLoader(sqlConfig *SQLConfig) *ConfigLoader {
	return &ConfigLoader{
		SQLConfig: sqlConfig,
	}
}

// Define the order of table creation
var tableOrder = []string{
	"users",
	"refresh",
	"llm",
	"integrations",
	"user_profile",
	"health",
	"conversations",
	"preferences",
	"inventory",
	"recipes",
}

func (loader *ConfigLoader) LoadSchemas() error {
	for _, tableName := range tableOrder {
		table, exists := Tables[tableName]
		if !exists {
			return fmt.Errorf("table %s not found in Tables map", tableName)
		}

		schema, exists := Schemas[tableName]
		if !exists {
			return fmt.Errorf("schema for table %s not found", tableName)
		}

		log.Printf("Creating table %s in database %s", table.Name, table.DbName)
		err := loader.SQLConfig.CreateTable(table.DbName, table.Name, schema)
		if err != nil {
			return fmt.Errorf("error creating table %s in database %s: %w", table.Name, table.DbName, err)
		}
		log.Printf("Table %s loaded successfully in database %s", table.Name, table.DbName)
	}
	return nil
}

func (loader *ConfigLoader) TestDatabases() error {
	for dbName := range loader.SQLConfig.Databases {
		if err := loader.testDatabase(dbName); err != nil {
			return fmt.Errorf("error testing database %s: %w", dbName, err)
		}
	}
	return nil
}

func (loader *ConfigLoader) testDatabase(dbName string) error {
	log.Printf("Testing database: %s", dbName)
	for _, tableName := range tableOrder {
		table := Tables[tableName]
		if table.DbName != dbName {
			continue
		}
		if err := loader.testTable(dbName, tableName); err != nil {
			return err
		}
	}
	return nil
}

func (loader *ConfigLoader) testTable(dbName, tableName string) error {
	log.Printf("Testing table: %s", tableName)

	// Test INSERT
	columns := Tables[tableName].Columns
	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = fmt.Sprintf("test_%s_%d", tableName, i)
	}

	_, err := loader.SQLConfig.Insert(dbName, tableName, columns, values...)
	if err != nil {
		return fmt.Errorf("error inserting into table %s: %w", tableName, err)
	}

	// Test SELECT
	query := fmt.Sprintf("SELECT * FROM %s LIMIT 1", tableName)
	rows, err := loader.SQLConfig.Query(dbName, query)
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
