package quiz

type Service interface {
	GenerateQuiz(GenrateQuizDTO *GenrateQuizDTO) (interface{}, error)
}
