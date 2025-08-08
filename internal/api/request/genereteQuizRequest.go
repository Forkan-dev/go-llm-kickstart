package request

import "github.com/gin-gonic/gin"

type GenrateQuizRequest struct {
	Subject    string `form:"subject" json:"subject" binding:"required"`
	Topic      string `form:"topic" json:"topic" binding:"required"`
	Difficulty string `form:"difficulty" json:"difficulty" binding:"required"`
	Type       string `form:"type" json:"type" binding:"required"`
	Format     string `form:"format" json:"format" binding:"required"`
}

func ValidateGenerateQuiz(c *gin.Context) (*GenrateQuizRequest, map[string]string) {
	return Validate(c, &GenrateQuizRequest{})
}
