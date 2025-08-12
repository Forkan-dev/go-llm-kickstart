package frontend

import (
	"learning-companion/internal/api/handlers/protected/frontend"
	"learning-companion/internal/config"
	"learning-companion/internal/service/llm"
	"learning-companion/internal/service/quiz"

	"github.com/gin-gonic/gin"
)

func NewQuizGenerateRoutes(router *gin.RouterGroup, cfg *config.Config) {
	llmService := llm.NewLLMService("llama3.2:latest")
	quizService := quiz.NewQuizService(llmService)
	quizHandler := frontend.NewQuizGeneratHandler(quizService)
	router.POST("/quiz-generate", quizHandler.GenerateQuiz) // Example protected route
}
