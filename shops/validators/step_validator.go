package validators

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/google/uuid"
)

// ValidateStep validates the required fields for a specific step
func ValidateStep(req *models.ValidateStepRequest) error {
	// First validate the step number
	if err := validate.Struct(req); err != nil {
		return err
	}

	// Get validation rules for the step
	rules := models.GetStepValidationRules()
	_, exists := rules[req.Step]
	if !exists {
		return fmt.Errorf("invalid step number: %d", req.Step)
	}

	// Get all fields that should have been validated in previous steps
	var allRequiredFields []string
	for i := 1; i <= req.Step; i++ {
		if stepRule, ok := rules[i]; ok {
			allRequiredFields = append(allRequiredFields, stepRule.RequiredFields...)
		}
	}

	// Validate required fields for the current step
	var missingFields []string

	for _, fieldName := range allRequiredFields {
		if isEmpty(req, fieldName) {
			missingFields = append(missingFields, fieldName)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("required fields missing: %s", strings.Join(missingFields, ", "))
	}

	// Additional validation for specific field types
	if req.Email != "" && !isValidEmail(req.Email) {
		return fmt.Errorf("invalid email format")
	}

	if req.Logo != "" && !isValidURL(req.Logo) {
		return fmt.Errorf("invalid logo URL format")
	}

	if req.Banner != "" && !isValidURL(req.Banner) {
		return fmt.Errorf("invalid banner URL format")
	}

	if req.OwnerID != uuid.Nil {
		if err := validate.Var(req.OwnerID, "uuid"); err != nil {
			return fmt.Errorf("invalid owner ID format")
		}
	}

	return nil
}

// isEmpty checks if a field in the struct is empty
func isEmpty(req *models.ValidateStepRequest, fieldName string) bool {
	v := reflect.ValueOf(*req)

	// Convert fieldName to match struct field names (capitalize first letter)
	structFieldName := strings.ToUpper(fieldName[:1]) + fieldName[1:]

	field := v.FieldByName(structFieldName)
	if !field.IsValid() {
		return true
	}

	switch field.Kind() {
	case reflect.String:
		return strings.TrimSpace(field.String()) == ""
	case reflect.Struct:
		// For UUID type
		if field.Type().String() == "uuid.UUID" {
			return field.Interface().(uuid.UUID) == uuid.Nil
		}
		return false
	default:
		return field.IsZero()
	}
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	err := validate.Var(email, "email")
	return err == nil
}

// isValidURL validates URL format
func isValidURL(url string) bool {
	err := validate.Var(url, "url")
	return err == nil
}
