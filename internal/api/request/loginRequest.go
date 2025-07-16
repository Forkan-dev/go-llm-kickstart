package request

import (
	"fmt"
	"learning-companion/internal/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username        string `form:"username" binding:"required_without=Email,omitempty"`
	Email           string `form:"email" binding:"required_without=Username,omitempty,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`
}

func Validate(c *gin.Context, passwordValidationConfig *config.PasswordValidationConfig) map[string]string {
	var req LoginRequest
	errors := make(map[string]string)

	if err := c.ShouldBind(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				fieldName := strings.ToLower(fieldErr.Field())
				switch fieldErr.Tag() {
				case "required":
					errors[fieldName] = fmt.Sprintf("The %s is required.", fieldName)
				case "required_without":
					errors[fieldName] = fmt.Sprintf("Either the %s or %s is required.", fieldName, fieldErr.Param())
				case "email":
					errors[fieldName] = fmt.Sprintf("The %s must be a valid email address.", fieldName)
				case "eqfield":
					errors[fieldName] = fmt.Sprintf("The %s must match the %s.", fieldName, fieldErr.Param())
				default:
					errors[fieldName] = fmt.Sprintf("Invalid value for %s: %s", fieldName, fieldErr.Error())
				}
			}
		} else {
			errors["form"] = "Invalid form data. Please check your input."
		}
		return errors
	}

	if len(req.Password) < passwordValidationConfig.MinLength {
		errors["password"] = fmt.Sprintf("Password must be at least %d characters long", passwordValidationConfig.MinLength)
	}

	if len(req.Password) > passwordValidationConfig.MaxLength {
		errors["password"] = fmt.Sprintf("Password must be no more than %d characters long", passwordValidationConfig.MaxLength)
	}

	return errors
}
