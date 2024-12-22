package serializer

import (
	"encoding/json"
	"fmt"
)

// Serializer interface defines the methods for serialization.
type Serializer interface {
	Serialize(interface{}) (map[string]interface{}, error)
	Deserialize(map[string]interface{}, interface{}) error
}

// BaseSerializer is the default implementation of Serializer.
type BaseSerializer struct {
	Fields []string // Included fields
}

// Serialize serializes a struct into a map with optional field filtering.
func (s *BaseSerializer) Serialize(data interface{}) (map[string]interface{}, error) {
	// Convierte la estructura en JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println("JSON Data:", string(jsonData)) // Debugging

	// Convierte el JSON en un mapa
	var result map[string]interface{}
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}
	fmt.Println("Mapa Deserializado:", result) // Debugging

	// Filtrar campos si es necesario
	if len(s.Fields) > 0 {
		filtered := make(map[string]interface{})
		for _, field := range s.Fields {
			if value, ok := result[field]; ok {
				filtered[field] = value
			}
		}
		fmt.Println("Campos Filtrados:", filtered) // Debugging
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
