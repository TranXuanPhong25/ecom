package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Có thể thêm custom validators nếu cần
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
