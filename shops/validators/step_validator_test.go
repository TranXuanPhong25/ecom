package validators

import (
	"testing"

	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/google/uuid"
)

func TestValidateStep_Step1(t *testing.T) {
	// Test step 1 with valid data
	req := &models.ValidateStepRequest{
		Step:         1,
		Name:         "My Shop",
		BusinessType: "individual",
	}

	err := ValidateStep(req)
	if err != nil {
		t.Errorf("Expected no error for valid step 1, got: %v", err)
	}
}

func TestValidateStep_Step1_MissingName(t *testing.T) {
	// Test step 1 with missing name
	req := &models.ValidateStepRequest{
		Step:         1,
		BusinessType: "individual",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for missing name in step 1, got none")
	}
}

func TestValidateStep_Step1_MissingBusinessType(t *testing.T) {
	// Test step 1 with missing business type
	req := &models.ValidateStepRequest{
		Step: 1,
		Name: "My Shop",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for missing businessType in step 1, got none")
	}
}

func TestValidateStep_Step2(t *testing.T) {
	// Test step 2 with valid data (should also validate step 1 fields)
	req := &models.ValidateStepRequest{
		Step:         2,
		Name:         "My Shop",
		BusinessType: "individual",
		Location:     "123 Main St",
		Email:        "test@example.com",
	}

	err := ValidateStep(req)
	if err != nil {
		t.Errorf("Expected no error for valid step 2, got: %v", err)
	}
}

func TestValidateStep_Step2_MissingLocation(t *testing.T) {
	// Test step 2 with missing location
	req := &models.ValidateStepRequest{
		Step:         2,
		Name:         "My Shop",
		BusinessType: "individual",
		Email:        "test@example.com",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for missing location in step 2, got none")
	}
}

func TestValidateStep_Step2_InvalidEmail(t *testing.T) {
	// Test step 2 with invalid email
	req := &models.ValidateStepRequest{
		Step:         2,
		Name:         "My Shop",
		BusinessType: "individual",
		Location:     "123 Main St",
		Email:        "invalid-email",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for invalid email in step 2, got none")
	}
}

func TestValidateStep_Step3(t *testing.T) {
	// Test step 3 with valid data
	req := &models.ValidateStepRequest{
		Step:         3,
		Name:         "My Shop",
		BusinessType: "individual",
		Location:     "123 Main St",
		Email:        "test@example.com",
		Logo:         "https://example.com/logo.png",
		Banner:       "https://example.com/banner.png",
	}

	err := ValidateStep(req)
	if err != nil {
		t.Errorf("Expected no error for valid step 3, got: %v", err)
	}
}

func TestValidateStep_Step3_InvalidLogoURL(t *testing.T) {
	// Test step 3 with invalid logo URL
	req := &models.ValidateStepRequest{
		Step:         3,
		Name:         "My Shop",
		BusinessType: "individual",
		Location:     "123 Main St",
		Email:        "test@example.com",
		Logo:         "not-a-url",
		Banner:       "https://example.com/banner.png",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for invalid logo URL in step 3, got none")
	}
}

func TestValidateStep_InvalidStepNumber(t *testing.T) {
	// Test with invalid step number
	req := &models.ValidateStepRequest{
		Step: 0,
		Name: "My Shop",
	}

	err := ValidateStep(req)
	if err == nil {
		t.Error("Expected error for invalid step number, got none")
	}
}

func TestValidateStep_WithOwnerID(t *testing.T) {
	// Test with valid owner ID
	ownerID := uuid.New()
	req := &models.ValidateStepRequest{
		Step:         1,
		Name:         "My Shop",
		BusinessType: "individual",
		OwnerID:      ownerID,
	}

	err := ValidateStep(req)
	if err != nil {
		t.Errorf("Expected no error with valid owner ID, got: %v", err)
	}
}
