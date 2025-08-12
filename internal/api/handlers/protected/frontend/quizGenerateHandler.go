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
	req, validationErrors := request.ValidateGenerateQuiz(c)
	log.Default().Println("this is validation error and req", validationErrors, req)

	if validationErrors != nil {
		response.ValidationError(c, "Validation failed", validationErrors, http.StatusBadRequest)
		return
	}
	log.Default().Println(req)

	res, err := h.quizService.GenerateQuiz("http://localhost:11434", quiz.QuizConfig{
		Subject:       req.Subject,
		Topics:        []string{req.Topic},
		Difficulty:    req.Difficulty,
		QuestionTypes: []string{req.Format},
		QuestionCount: 10,
		Strategy:      "extreme",
		MaxTokens:     120,
	})
	// Generate the quiz
	if err != nil {
		log.Printf("Failed to generate quiz: %v", err)
		response.Error(c, "Failed to generate quiz", http.StatusInternalServerError)
		return
	}

	apiResponse := map[string]interface{}{
		"quiz":   res,
		"format": req.Format,
	}
	// Return the quiz
	response.Success(c, "Quiz generated successfully", apiResponse, http.StatusOK)

}
