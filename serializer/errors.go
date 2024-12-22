package serializer

import "fmt"

// ValidationError represents an error that occurred during validation.
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s (value: %v)", e.Field, e.Message, e.Value)
}

// TransformationError represents an error that occurred during a transformation.
type TransformationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *TransformationError) Error() string {
	return fmt.Sprintf("Transformation error on field '%s': %s (value: %v)", e.Field, e.Message, e.Value)
}

// SerializationError represents an error that occurred during serialization or deserialization.
type SerializationError struct {
	Message string
}

func (e *SerializationError) Error() string {
	return fmt.Sprintf("Serialization error: %s", e.Message)
}
