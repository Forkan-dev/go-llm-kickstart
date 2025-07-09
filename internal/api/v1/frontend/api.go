package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learning-companion/internal/api/middleware"
	"learning-companion/internal/response"
)

func RegisterRoutes(router *gin.RouterGroup, jwtSecret string) {
	// Public routes
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Frontend API v1 is up and running!",
		})
	})

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware(jwtSecret))
	{
		protected.GET("/protected", func(c *gin.Context) {
			response.Success(c, "This is a protected frontend  and working route!", nil, http.StatusOK)
		})
	}
}
