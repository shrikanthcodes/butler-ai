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

var tableOrder = []string{
	"users",
	"refresh",
	"user_profile",
	// Add other tables in the correct order...
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
