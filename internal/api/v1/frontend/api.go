package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learning-companion/internal/api/handlers/auth"
	"learning-companion/internal/api/middleware"
	"learning-companion/internal/config"
)

func RegisterRoutes(router *gin.RouterGroup, cfg *config.Config) {
	// Public routes
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Frontend API v1 is up and running!",
		})
	})

	// public routes
	public := router.Group("/")
	public.Use(middleware.JWTPublicMiddleware(cfg.Server.JWTSecret))
	{
		public.POST("/login", func(c *gin.Context) {
			auth.Login(c)
		})
	}
	// Protected routes
}