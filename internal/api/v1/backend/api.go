package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learning-companion/internal/api/middleware"
)

func RegisterRoutes(router *gin.RouterGroup, jwtSecret string) {
	backendGroup := router.Group("/api/v1/backend")

	// Public routes
	backendGroup.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend API v1 is up and running!",
		})
	})

	// Protected routes
	protected := backendGroup.Group("/")
	protected.Use(middleware.JWTAuthMiddleware(jwtSecret))
	{
		protected.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "This is a protected backend (admin) route!"})
		})
	}
}
