package serializer

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"gopkg.in/yaml.v3" // YAML library, install using: go get gopkg.in/yaml.v3
)

// Custom error types for better error handling

// Serializer interface defines the methods for serialization.
type Serializer interface {
	Serialize(interface{}) (map[string]interface{}, error)
	Deserialize(map[string]interface{}, interface{}) error
	Validate(map[string]interface{}) error
	SerializeToXML(interface{}) (string, error)
	SerializeToYAML(interface{}) (string, error)
}

// BaseSerializer is the default implementation of Serializer.
type BaseSerializer struct {
	Fields            []string                                     // Included fields
	Validations       map[string][]func(interface{}) error         // Multiple validations per field
	Transformations   map[string]func(interface{}) interface{}     // Transformations by field
	ConditionalFields map[string]func(map[string]interface{}) bool // Conditional inclusion of fields
}

// Serialize serializes a struct into a map with optional field filtering, transformations, and conditional fields.
func (s *BaseSerializer) Serialize(data interface{}) (map[string]interface{}, error) {
	// Convert struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, &SerializationError{Message: fmt.Sprintf("failed to serialize struct: %v", err)}
	}

	// Convert JSON to a map
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, &SerializationError{Message: fmt.Sprintf("failed to convert JSON to map: %v", err)}
	}

	// Apply transformations
	if s.Transformations != nil {
		for field, transform := range s.Transformations {
			if value, exists := result[field]; exists {
				transformedValue := transform(value)
				if transformedValue == nil {
					return nil, &TransformationError{
						Field:   field,
						Value:   value,
						Message: "transformation returned nil",
					}
				}
				result[field] = transformedValue
			}
		}
	}

	// Apply conditional fields
	if s.ConditionalFields != nil {
		for field, condition := range s.ConditionalFields {
			if include := condition(result); !include {
				delete(result, field) // Exclude the field if condition is false
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

// SerializeToXML serializes a struct into an XML string.
func (s *BaseSerializer) SerializeToXML(data interface{}) (string, error) {
	xmlData, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", &SerializationError{Message: fmt.Sprintf("failed to serialize to XML: %v", err)}
	}
	return string(xmlData), nil
}

// SerializeToYAML serializes a struct into a YAML string.
func (s *BaseSerializer) SerializeToYAML(data interface{}) (string, error) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", &SerializationError{Message: fmt.Sprintf("failed to serialize to YAML: %v", err)}
	}
	return string(yamlData), nil
}

// Deserialize deserializes a map into a struct.
func (s *BaseSerializer) Deserialize(input map[string]interface{}, out interface{}) error {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return &SerializationError{Message: fmt.Sprintf("failed to convert map to JSON: %v", err)}
	}
	if err := json.Unmarshal(jsonData, out); err != nil {
		return &SerializationError{Message: fmt.Sprintf("failed to deserialize JSON to struct: %v", err)}
	}
	return nil
}

// Validate checks the provided data against the validations defined in the serializer.
func (s *BaseSerializer) Validate(data map[string]interface{}) error {
	if s.Validations == nil {
		return nil // No validations defined
	}

	for field, validations := range s.Validations {
		if value, exists := data[field]; exists {
			for _, validation := range validations {
				if err := validation(value); err != nil {
					return &ValidationError{
						Field:   field,
						Value:   value,
						Message: err.Error(),
					}
				}
			}
		} else {
			return &ValidationError{
				Field:   field,
				Message: "field is missing",
			}
		}
	}

	return nil
}
