package enum

import (
	"errors"
	"fmt"
	"reflect"
)

// generateEnumMap generates a map of valid values for a given enum type.
func generateEnumMap(enumType interface{}) (map[string]bool, error) {
	result := make(map[string]bool)

	enumValue := reflect.TypeOf(enumType)
	if enumValue.Kind() != reflect.Int {
		return nil, errors.New("provided type is not a valid enum")
	}

	for i := 0; i < reflect.TypeOf(enumType).NumMethod(); i++ {
		name := enumType.(fmt.Stringer).String()
		result[name] = true
	}
	return result, nil
}

// Setter checks if a given value is valid for the specified enum type.
func Setter(enumName, value string) error {
	enumType, ok := enumTypes[enumName]
	if !ok {
		return fmt.Errorf("enum type '%s' not found", enumName)
	}

	validValues, err := generateEnumMap(enumType)
	if err != nil {
		return fmt.Errorf("failed to generate enum map for '%s': %w", enumName, err)
	}

	if !validValues[value] {
		return fmt.Errorf("invalid value '%s' for enum type '%s'", value, enumName)
	}
	return nil
}
