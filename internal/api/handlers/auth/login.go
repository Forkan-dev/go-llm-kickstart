package auth

import (
	"learning-companion/internal/api/request"
	"learning-companion/internal/model"
	"learning-companion/internal/response"
	"learning-companion/pkg/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Token     string `json:"access_token"`
	ExpiresAt string `json:"expires_at"`
}

func Login(c *gin.Context) {
	errors := request.Validate(c)

	if len(errors) > 0 {
		response.ValidationError(c, "Validation failed", errors, http.StatusBadRequest)
		return
	}

	user := model.User{}
	// get the user form database use ORM\

	database.DB.Model(&user).Where("username = ? OR email = ?", c.PostForm("username"), c.PostForm("email")).First(&user)

	// Check if user exists
	if user.ID == 0 {
		response.Error(c, "User not found", http.StatusNotFound)
		return
	}

	// Check bycript password

	if !user.CheckPassword(c.PostForm("password")) {
		response.Error(c, "Invalid password", http.StatusUnauthorized)
		return
	}

	// For example, validate user credentials and generate JWT token
	acessToken := user.CreateToken()
	if acessToken == "" {
		response.Error(c, "Failed to create access token", http.StatusInternalServerError)
		return
	}

	parseToken, err := user.ParseToken(acessToken)
	if err != nil {
		response.Error(c, "Failed to parse access token", http.StatusInternalServerError)
		return
	}

	loginResponse := LoginResponse{
		Token:     acessToken,
		ExpiresAt: parseToken.ExpiresAt.Time.Format(time.RFC3339),
	}

	response.Success(c, "Login successful", loginResponse, http.StatusOK)
}
