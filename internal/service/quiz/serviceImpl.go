package quiz

import "log"

type GenrateQuizDTO struct {
	Subject    string `json:"subject"`
	Topic      string `json:"topic"`
	Difficulty string `json:"difficulty"`
	Type       string `json:"type"`
	Format     string `json:"format"`
}

type quizImplService struct {
}

func NewService() Service {
	return &quizImplService{}
}

func (s *quizImplService) GenerateQuiz(GenrateQuizDTO *GenrateQuizDTO) (string, error) {
	log.Default().Println(GenrateQuizDTO)
	return "", nil
}
