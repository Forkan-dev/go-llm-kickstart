package validator

import (
	"fmt"
	"learning-companion/internal/config"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var PasswordValidationConfig *config.PasswordValidationConfig

// NewPasswordValidator is a factory that creates our validation function.
// It "closes over" the passwordValidationConfig to make it available at runtime.
func NewPasswordValidator(cfg *config.PasswordValidationConfig) validator.Func {
	return func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		//set min man to fl
		if len(password) < cfg.MinLength || len(password) > cfg.MaxLength {
			return false // Validation fails
		}
		return true // Validation passes
	}
}

// GetErrorMsg returns a more specific error message for validation failures.
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "The " + fe.Field() + " field is required"
	case "email":
		return "Invalid email format"
	case "eqfield":
		return "The field " + fe.Field() + " must match the " + fe.Param() + " field"
	case "password":
		return fmt.Sprintf("Password must be between %d and %d characters long", PasswordValidationConfig.MinLength, PasswordValidationConfig.MaxLength)
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
