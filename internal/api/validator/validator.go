package validator

import (
	"fmt"
	"learning-companion/internal/config"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// NewPasswordValidator is a factory that creates our validation function.
// It "closes over" the passwordValidationConfig to make it available at runtime.
func NewPasswordValidator(cfg *config.PasswordValidationConfig) validator.Func {
	return func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		fmt.Printf("DEBUG: Validating password. Length: %d, MinLength: %d, MaxLength: %d\n", len(password), cfg.MinLength, cfg.MaxLength)
		if len(password) < cfg.MinLength || len(password) > cfg.MaxLength {
			fmt.Println("DEBUG: Password validation failed.")
			return false // Validation fails
		}
		fmt.Println("DEBUG: Password validation passed.")
		return true // Validation passes
	}
}

// GetErrorMsg returns a more specific error message for validation failures.
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "eqfield":
		return "The field " + fe.Field() + " must match the " + fe.Param() + " field"
	case "password":
		return "Password must be between 8 and 64 characters long"
	case "required_without":
		return "This field is required if " + fe.Param() + " is not present"
	}
	return "Unknown validation error"
}

// GetFieldName returns the lowercase JSON field name.
func GetFieldName(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
