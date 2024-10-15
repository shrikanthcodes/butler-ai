package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONB is a generic wrapper that implements driver.Valuer and sql.Scanner.
// It can be used to store and retrieve JSONB fields from the database.
type JSONB[T any] struct {
	Data T
}

// Value implements the driver.Valuer interface to convert Go struct/slice into JSON string.
func (j *JSONB[T]) Value() (driver.Value, error) {
	value, err := json.Marshal(j.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSONB: %w", err)
	}
	return string(value), nil
}

// Scan implements the sql.Scanner interface to convert JSON string into Go struct/slice.
func (j *JSONB[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB: %v", value)
	}
	err := json.Unmarshal(bytes, &j.Data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB: %w", err)
	}
	return nil
}
