package request

import (
	"learning-companion/internal/api/validator"
	"log"
	"reflect"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	validator10 "github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username        string `form:"username" json:"username" binding:"required_without=Email,omitempty"`
	Email           string `form:"email" json:"email" binding:"required_without=Username,omitempty,email"`
	Password        string `form:"password" json:"password" binding:"required,password"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}

func Validate(c *gin.Context) (*LoginRequest, map[string]string) {
	var req LoginRequest
	errors := make(map[string]string)
	reqType := reflect.TypeOf(req)

	// Bind and validate the request
	if err := c.ShouldBindJSON(&req); err != nil {
		if validationErrors, ok := err.(validator10.ValidationErrors); ok {
			// Map validation errors to JSON field names
			for _, fieldErr := range validationErrors {
				field, found := reqType.FieldByName(fieldErr.Field())
				if !found {
					log.Printf("Field %s not found in LoginRequest struct: %s", fieldErr.Field(), debug.Stack())
					continue
				}
				jsonTag := validator.GetFieldName(field)
				if jsonTag != "" {
					errors[jsonTag] = validator.GetErrorMsg(fieldErr)
				}
			}
		} else {
			// Handle non-validation errors (e.g., malformed JSON)
			errors["request"] = "Invalid request body. Please check your input."
		}
		return nil, errors
	}

	// Return the validated request and no errors
	return &req, nil
}
