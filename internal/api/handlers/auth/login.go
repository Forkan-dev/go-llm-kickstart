package auth

import (
	"context"
	"encoding/json"
	"learning-companion/internal/api/request"
	"learning-companion/internal/response"
	"learning-companion/internal/service/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/ollama"
)

type LoginResponse struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}

var authService = auth.NewService()

func Login(c *gin.Context) {
	// Validate the request and get the request data
	req, errors := request.Validate(c)
	if errors != nil {
		response.ValidationError(c, "Validation failed", errors, http.StatusBadRequest)
		return
	}

	if req.Username == "" {
		req.Username = req.Email
	}

	// Proceed with the login logic
	user, accessToken, refreshTokenString, err := authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, err.Error(), http.StatusUnauthorized)
		return
	}

	parsedToken, err := user.ParseToken(accessToken)
	if err != nil {
		response.Error(c, "Failed to parse access token", http.StatusInternalServerError)
		return
	}

	loginResponse := LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshTokenString,
		ExpiresAt:    parsedToken.ExpiresAt.Time.Format(time.RFC3339),
	}

	response.Success(c, "Login successful", loginResponse, http.StatusOK)
}

func Logout(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	err := authService.Logout(accessToken)
	if err != nil {
		response.Error(c, err.Error(), http.StatusUnauthorized)
		return
	}

	response.Success(c, "Logout successful", nil, http.StatusOK)
}

type AiTestRequest struct {
	Prompt string `json:"prompt" binding:"required"`
}

func Aitesting(c *gin.Context) {
	// Parse JSON body
	prompt := c.Query("prompt")
	if prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Prompt is required"})
		return
	}

	// Initialize Ollama model
	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Printf("Ollama init error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI model initialization failed"})
		return
	}

	ctx := context.Background()

	// Generate response from Ollama
	resp, err := llm.Call(ctx, prompt)
	//json decode from string resp
	var jsonResp map[string]interface{}
	if err := json.Unmarshal([]byte(resp), &jsonResp); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI response parsing failed"})
		return
	}
	if err != nil {
		log.Printf("Ollama call error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI generation failed"})
		return
	}

	resp2, err := llm.Call(ctx, prompt+" add 444 and multiply by 2 Respond with only one number as JSON: {"+"word"+": "+"value"+"}")
	// Return success response with AI output
	response.Success(c, "AI response generated", gin.H{"result": jsonResp, "result2": resp2}, http.StatusOK)
}

// func Login(c *gin.Context) {
// 	var req request.LoginRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		if errs, ok := err.(validator.ValidationErrors); ok {
// 			errors := make(map[string]string)
// 			reqType := reflect.TypeOf(req)

// 			for _, e := range errs {
// 				field, _ := reqType.FieldByName(e.Field())
// 				jsonTag := apivalidator.GetFieldName(field)
// 				if jsonTag != "" {
// 					errors[jsonTag] = apivalidator.GetErrorMsg(e)
// 				}
// 			}
// 			response.ValidationError(c, "Validation failed", errors, http.StatusBadRequest)
// 			return
// 		}

// 		// Handle other errors (e.g., invalid JSON)
// 		response.Error(c, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	user, accessToken, refreshTokenString, err := authService.Login(req.Username, req.Password)
// 	if err != nil {
// 		response.Error(c, err.Error(), http.StatusUnauthorized)
// 		return
// 	}

// 	parsedToken, err := user.ParseToken(accessToken)
// 	if err != nil {
// 		response.Error(c, "Failed to parse access token", http.StatusInternalServerError)
// 		return
// 	}

// 	loginResponse := LoginResponse{
// 		Token:        accessToken,
// 		RefreshToken: refreshTokenString,
// 		ExpiresAt:    parsedToken.ExpiresAt.Time.Format(time.RFC3339),
// 	}

// 	response.Success(c, "Login successful", loginResponse, http.StatusOK)
// }
