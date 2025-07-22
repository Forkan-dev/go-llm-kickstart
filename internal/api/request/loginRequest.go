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

func Validate(c *gin.Context) map[string]string {
	var req LoginRequest
	errors := make(map[string]string)
	reqType := reflect.TypeOf(req)

	if err := c.ShouldBind(&req); err != nil {
		if validationErrors, ok := err.(validator10.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				field, found := reqType.FieldByName(fieldErr.Field())
				if !found {
					log.Printf("Field %s not found in LoginRequest struct: %s", fieldErr.Field(), debug.Stack())
					continue
				}

				jsonTag := field.Tag.Get("json")
				errors[jsonTag] = validator.GetErrorMsg(fieldErr)
			}
		} else {
			errors["form"] = "Invalid form data. Please check your input."
		}
	}
	return errors
}
