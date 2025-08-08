package frontend

import (
	"learning-companion/internal/api/request"
	"learning-companion/internal/response"
	"learning-companion/internal/service/quiz"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuizGenerateHandler struct {
	quizService quiz.Service
}

func NewQuizGeneratHandler(service quiz.Service) *QuizGenerateHandler {
	return &QuizGenerateHandler{
		quizService: service,
	}
}
func (h *QuizGenerateHandler) GenerateQuiz(c *gin.Context) {
	// Get the request body
	req, err := request.ValidateGenerateQuiz(c)

	if err != nil {
		response.ValidationError(c, "Validation failed", err, http.StatusBadRequest)
		return
	}
	log.Default().Println(req)

	h.quizService.GenerateQuiz(&quiz.GenrateQuizDTO{
		Subject:    req.Subject,
		Topic:      req.Topic,
		Difficulty: req.Difficulty,
		Type:       req.Type,
		Format:     req.Format,
	})
	// Generate the quiz
	// Return the quiz
	response.Success(c, "Quiz generated successfully", nil, http.StatusOK)

}
