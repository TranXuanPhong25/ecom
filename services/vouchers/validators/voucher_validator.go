package validators

import (
	"github.com/go-playground/validator/v10"
)

// NewValidator creates a new validator instance
func NewValidator() *validator.Validate {
	return validator.New()
}
