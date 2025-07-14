package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// Có thể thêm custom validators nếu cần
	_ = validate.RegisterValidation("strong_password", validateStrongPassword)
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return len(password) >= 8 && hasUpper && hasLower && hasNumber && hasSpecial
}
