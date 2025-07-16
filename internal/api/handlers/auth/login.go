package auth

import (
	"learning-companion/internal/api/request"
	"learning-companion/internal/config"
	"learning-companion/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, passwordValidationConfig *config.PasswordValidationConfig) {
	errors := request.Validate(c, passwordValidationConfig)

	if len(errors) > 0 {
		response.ValidationError(c, "Validation failed", errors, http.StatusBadRequest)
		return
	}
	// Handle login logic here
	// For example, validate user credentials and generate JWT token

	response.Success(c, "Login successful", nil, http.StatusOK)
}
