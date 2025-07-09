package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/status/backend", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend API v1 is up and running!",
		})
	})
}
