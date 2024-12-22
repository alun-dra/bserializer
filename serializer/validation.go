package serializer

import (
	"fmt"
	"strings"
)

// NotEmpty checks if a field is not empty.
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

// Positive checks if a field is a positive number.
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

// ValidPassword checks if a password meets certain criteria.
func ValidPassword(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}
	if len(str) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !strings.ContainsAny(str, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !strings.ContainsAny(str, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !strings.ContainsAny(str, "0123456789") {
		return fmt.Errorf("password must contain at least one number")
	}
	if !strings.ContainsAny(str, "!@#$%^&*()_+=-") {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}
