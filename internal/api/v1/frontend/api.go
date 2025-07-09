package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ 
			"message": "Frontend API v1 is up and running!",
		})
	})
}