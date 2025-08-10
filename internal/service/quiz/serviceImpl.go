package quiz

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/ollama"
)

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

type TF struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type MCQ struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
}

type Question struct {
	TF
}

const (
	TFFormat  = "True/False"
	MCQFormat = "MCQ"
)

func buildQuizQuestionPropmt(GenrateQuizDTO *GenrateQuizDTO, quantity int) string {

	subject := GenrateQuizDTO.Subject
	topic := GenrateQuizDTO.Topic
	if subject == "" && topic == "" {
		subject = "General Knowledge"
	}

	return fmt.Sprintf(
		"Generate exactly %d short questions and answer for %s about %s with %s difficulty. "+
			"Only output the list of questions, no answers, explanations, or extra text. "+
			"Each question must be concise, clear, and directly related to the topic.",
		quantity,
		subject,
		topic,
		GenrateQuizDTO.Difficulty,
	)
}

func buildQuizFormatPrompt(GenrateQuizDTO *GenrateQuizDTO, questionsPrompt string, quantity int) (string, string) {

	format := GenrateQuizDTO.Format
	switch GenrateQuizDTO.Format {
	case "mcq", "MCQ", "Mcq":
		format = "MCQ"
	case "TF", "True/False", "true/false", "truefalse", "TrueFalse":
		format = "True/False"
	default:
		format = "Question"
	}

	return format, fmt.Sprintf(
		"Fist make sure the quantity is %d.Format the following questions and answers into strict %s JSON format. "+
			"Questions and answers: %s. "+
			"If the format is MCQ, each item must have exactly these keys: "+
			"\"question\" (string), \"options\" (array of strings), and \"answer\" (string, exactly matching one of the provided options). "+
			"If the format is True/False, each item must have \"question\" (string) and \"answer\" (boolean). "+
			"If the format is Short Answer or Open ended, each item must have \"question\" (string) and \"answer\" (string). "+
			"Return ONLY a valid JSON array of quiz objects. "+
			"No text outside the JSON. No explanations. No comments. The JSON must be syntactically valid and ready for parsing.",
		quantity,
		format,
		questionsPrompt,
	)
}

func (s *quizImplService) GenerateQuiz(GenrateQuizDTO *GenrateQuizDTO) (interface{}, error) {

	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Printf("Ollama init error: %v", err)
		return nil, err
	}
	prompt := buildQuizQuestionPropmt(GenrateQuizDTO, 5)
	ctx := context.Background()
	resp, err := llm.Call(ctx, prompt)
	if err != nil {
		log.Printf("Ollama call error: %v", err)
		return nil, err
	}

	format, formatPrompt := buildQuizFormatPrompt(GenrateQuizDTO, resp, 5)
	resp, err = llm.Call(ctx, formatPrompt)
	if err != nil {
		log.Printf("Ollama call error: %v", err)
		return nil, err
	}
	log.Default().Println("resp -----==========>:", resp)
	//json decode from string resp
	var formatQ interface{}
	switch format {
	case TFFormat:
		formatQ = &[]TF{}
	case MCQFormat:
		formatQ = &[]MCQ{}
	default:
		formatQ = &[]Question{}
	}

	if err := json.Unmarshal([]byte(resp), &formatQ); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return nil, err
	}

	log.Default().Println("jsonResp:", formatQ)

	return formatQ, nil
}
