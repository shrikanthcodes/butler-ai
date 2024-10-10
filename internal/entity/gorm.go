package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// DialogueArray defines the dialogue array to store the transcript(for jsonb field)
type DialogueArray []Dialogue

// Value implements the driver.Valuer interface to convert Go struct to JSON for storing in DB.
func (d *DialogueArray) Value() (driver.Value, error) {
	value, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal DialogueArray: %w", err)
	}
	return string(value), nil
}

// Scan implements the sql.Scanner interface to scan JSON from the DB into Go struct.
func (d *DialogueArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal DialogueArray: %v", value)
	}

	err := json.Unmarshal(bytes, d)
	if err != nil {
		return fmt.Errorf("failed to unmarshal DialogueArray: %w", err)
	}
	return nil
}
