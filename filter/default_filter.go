package filter

import "strings"

// Validator is an implementation of the Filter interface.
type Validator struct{}

// NewValidator creates a new Validator.
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks if the message is valid.
func (v *Validator) Validate(data string) bool {
	return strings.Contains(data, "valid")
}
