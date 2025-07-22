package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learning-companion/internal/api/handlers/auth"
	"learning-companion/internal/api/middleware"
	"learning-companion/internal/config"
	"learning-companion/internal/response"
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
		public.POST("/login", auth.Login)
	}
	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			response.Success(c, "This is a protected route", nil, http.StatusOK)
		}) // Example protected route

		protected.POST("/logout", auth.Logout) // Example logout route
	}
}
