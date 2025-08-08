package request

import (
	"learning-companion/internal/api/validator"
	"log"
	"reflect"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	validator10 "github.com/go-playground/validator/v10"
)

func Validate[T any](c *gin.Context, validationReq *T) (*T, map[string]string) {

	errors := make(map[string]string)
	reqType := reflect.TypeOf(validationReq).Elem()

	// Bind and validate the request
	if err := c.ShouldBindJSON(&validationReq); err != nil {
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
	return validationReq, nil
}
