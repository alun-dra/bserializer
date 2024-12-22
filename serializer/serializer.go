package serializer

import (
	"encoding/json"
	"fmt"
)

// Serializer interface defines the methods for serialization.
type Serializer interface {
	Serialize(interface{}) (map[string]interface{}, error)
	Deserialize(map[string]interface{}, interface{}) error
	Validate(map[string]interface{}) error
}

// BaseSerializer is the default implementation of Serializer.
type BaseSerializer struct {
	Fields          []string                                 // Included fields
	Validations     map[string]func(interface{}) error       // Validations by field
	Transformations map[string]func(interface{}) interface{} // Transformations by field
}

// Serialize serializes a struct into a map with optional field filtering and transformations.
func (s *BaseSerializer) Serialize(data interface{}) (map[string]interface{}, error) {
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Convert JSON to a map
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}

	// Apply transformations
	if s.Transformations != nil {
		for field, transform := range s.Transformations {
			if value, exists := result[field]; exists {
				result[field] = transform(value)
			}
		}
	}

	// Filter fields if necessary
	if len(s.Fields) > 0 {
		filtered := make(map[string]interface{})
		for _, field := range s.Fields {
			if value, ok := result[field]; ok {
				filtered[field] = value
			} else {
				filtered[field] = nil // Default to nil if field is missing
			}
		}
		return filtered, nil
	}

	return result, nil
}

// Deserialize deserializes a map into a struct.
func (s *BaseSerializer) Deserialize(input map[string]interface{}, out interface{}) error {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, out)
}

// Validate checks the provided data against the validations defined in the serializer.
func (s *BaseSerializer) Validate(data map[string]interface{}) error {
	if s.Validations == nil {
		return nil // No validations defined
	}

	for field, validation := range s.Validations {
		if value, exists := data[field]; exists {
			if err := validation(value); err != nil {
				return fmt.Errorf("validation failed for field '%s': %w", field, err)
			}
		} else {
			return fmt.Errorf("field '%s' is missing", field)
		}
	}

	return nil
}
