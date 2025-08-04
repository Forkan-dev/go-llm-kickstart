package frontend

import (
	"learning-companion/internal/api/handlers/protected/frontend"
	"learning-companion/internal/config"
	"learning-companion/internal/service/subject"
	"learning-companion/pkg/database"

	"github.com/gin-gonic/gin"
)

func NewSubjectRoutes(router *gin.RouterGroup, cfg *config.Config) {
	// Register the subject routes
	subjectService := subject.NewService(database.DB)
	subjectHandler := frontend.NewSubjectHandler(subjectService)
	router.GET("/subjects", subjectHandler.GetSubjects) // Example protected route
}
