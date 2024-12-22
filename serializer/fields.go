package serializer

import "fmt"

// Field interface for validating field values.
type Field interface {
	Validate(value interface{}) error
}

// StringField represents a string with validation options.
type StringField struct {
	MaxLength int
}

func (f StringField) Validate(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("not a valid string")
	}
	if len(str) > f.MaxLength {
		return fmt.Errorf("string exceeds max length")
	}
	return nil
}
