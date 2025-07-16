package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learning-companion/internal/api/middleware"
	"learning-companion/internal/config"
	"learning-companion/internal/response"
)

func RegisterRoutes(router *gin.RouterGroup, cfg *config.Config) {
	backendGroup := router.Group("/api/v1/backend")

	// Public routes
	backendGroup.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend API v1 is up and running!",
		})
	})

	// public routes
	public := backendGroup.Group("/")
	public.Use(middleware.JWTPublicMiddleware(cfg.Server.JWTSecret))
	{
		public.GET("/protected", func(c *gin.Context) {
			response.Success(c, "This is a public Backend and working route!", nil, http.StatusOK)
		})
	}
}
