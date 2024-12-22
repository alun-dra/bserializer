package serializer

import "fmt"

// NotEmpty is a predefined validation function that checks if a field is not empty.
func NotEmpty(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}
	if str == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}

// Positive is a predefined validation function that checks if a field is a positive number.
func Positive(value interface{}) error {
	num, ok := value.(float64) // JSON numbers are parsed as float64
	if !ok {
		return fmt.Errorf("value is not a number")
	}
	if num <= 0 {
		return fmt.Errorf("value must be positive")
	}
	return nil
}
