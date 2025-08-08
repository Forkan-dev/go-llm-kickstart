package quiz

type Service interface {
	GenerateQuiz(GenrateQuizDTO *GenrateQuizDTO) (string, error)
}
