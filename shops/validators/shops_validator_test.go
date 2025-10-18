package validators

import (
	"testing"

	"github.com/TranXuanPhong25/ecom/shops/models"
	"github.com/google/uuid"
)

func TestCreateShopValidation(t *testing.T) {
	tests := []struct {
		name        string
		request     models.CreateShopRequest
		shouldError bool
		errorCount  int
	}{
		{
			name: "Valid request with all fields",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Logo:         "https://example.com/logo.png",
				Banner:       "https://example.com/banner.png",
				Email:        "test@example.com",
				Phone:        "1234567890",
				BusinessType: "individual",
			},
			shouldError: false,
		},
		{
			name: "Valid request with empty phone",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "test@example.com",
				Phone:        "",
				BusinessType: "business",
			},
			shouldError: false,
		},
		{
			name: "Invalid phone with letters",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "test@example.com",
				Phone:        "abc123",
				BusinessType: "individual",
			},
			shouldError: true,
		},
		{
			name: "Invalid business type",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "test@example.com",
				BusinessType: "invalid",
			},
			shouldError: true,
		},
		{
			name: "Missing required name",
			request: models.CreateShopRequest{
				Name:         "",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "test@example.com",
				BusinessType: "individual",
			},
			shouldError: true,
		},
		{
			name: "Missing required email",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "",
				BusinessType: "individual",
			},
			shouldError: true,
		},
		{
			name: "Invalid email format",
			request: models.CreateShopRequest{
				Name:         "Test Shop",
				OwnerID:      uuid.New(),
				Location:     "123 Main St",
				Email:        "notanemail",
				BusinessType: "individual",
			},
			shouldError: true,
		},
		{
			name: "Missing required business type",
			request: models.CreateShopRequest{
				Name:     "Test Shop",
				OwnerID:  uuid.New(),
				Location: "123 Main St",
				Email:    "test@example.com",
			},
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStruct(&tt.request)
			if tt.shouldError && err == nil {
				t.Errorf("Expected validation error but got none")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("Expected no validation error but got: %v", err)
			}
			if err != nil {
				errors := FormatValidationErrors(err)
				t.Logf("Validation errors: %+v", errors)
			}
		})
	}
}
