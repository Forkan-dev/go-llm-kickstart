package quiz

type Service interface {
	GenerateQuiz(baseURL string, config QuizConfig) (*QuizResponse, error)
	ValidateQuizConfig(config QuizConfig) error
	ParseQuizResponse(llmResponse string) (*QuizResponse, error)
	BuildPrompt(config QuizConfig) (string, error)
}
