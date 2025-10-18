package models

import "github.com/google/uuid"

// ValidateStepRequest represents a request to validate a specific step in the shop creation process
type ValidateStepRequest struct {
	Step         int       `json:"step" validate:"required,min=1,max=3"`
	Name         string    `json:"name"`
	OwnerID      uuid.UUID `json:"ownerId"`
	Location     string    `json:"location"`
	Logo         string    `json:"logo"`
	Banner       string    `json:"banner"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	BusinessType string    `json:"businessType"`
}

// StepValidationRules defines validation rules for each step
type StepValidationRules struct {
	RequiredFields []string
}

// GetStepValidationRules returns the validation rules for each step
func GetStepValidationRules() map[int]StepValidationRules {
	return map[int]StepValidationRules{
		1: {
			RequiredFields: []string{"name", "businessType"},
		},
		2: {
			RequiredFields: []string{"location", "email"},
		},
		3: {
			RequiredFields: []string{"logo", "banner"},
		},
	}
}
