package validators

import (
	"github.com/TranXuanPhong25/ecom/services/product-reviews/dtos"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateReview validates the CreateReviewRequest
func ValidateCreateReview(req *dtos.CreateReviewRequest) error {
	return validate.Struct(req)
}

// ValidateUpdateReview validates the UpdateReviewRequest
func ValidateUpdateReview(req *dtos.UpdateReviewRequest) error {
	return validate.Struct(req)
}

// GetValidationErrors formats validation errors into a readable string
func GetValidationErrors(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				return e.Field() + " is required"
			case "min":
				return e.Field() + " must be at least " + e.Param()
			case "max":
				return e.Field() + " must be at most " + e.Param()
			}
		}
	}
	return "Validation failed"
}
