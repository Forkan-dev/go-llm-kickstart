package response

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type ValidationErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Status  int         `json:"status"`
}

func Success(c *gin.Context, message string, data interface{}, status int) {
	response := SuccessResponse{
		Message: message,
		Data:    data,
		Status:  status,
	}

	c.JSON(status, response)
}

func Error(c *gin.Context, message string, status int) {
	response := ErrorResponse{
		Message: message,
		Status:  status,
	}

	c.JSON(status, response)
}

func ValidationError(c *gin.Context, message string, errors interface{}, status int) {
	response := ValidationErrorResponse{
		Message: message,
		Errors:  errors,
		Status:  status,
	}

	c.JSON(status, response)
}

func NotFound(c *gin.Context, message string) {
	Error(c, message, 404)
}
