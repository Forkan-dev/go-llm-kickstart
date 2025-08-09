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
	MCQ MCQ `json:"mcq"`
}

const (
	TFFormat  = "True/False"
	MCQFormat = "MCQ"
)

func buildQuizPrompt(GenrateQuizDTO *GenrateQuizDTO) (string, string) {
	format := GenrateQuizDTO.Format
	switch GenrateQuizDTO.Format {
	case "mcq", "MCQ", "Mcq":
		format = "MCQ"
	case "TF", "True/False", "true/false", "truefalse", "TrueFalse":
		format = "True/False"
	default:
		format = "Question"
	}

	subject := GenrateQuizDTO.Subject
	topic := GenrateQuizDTO.Topic
	if subject == "" && topic == "" {
		subject = "General Knowledge"
	}

	return format, fmt.Sprintf(`Generate a quiz based on the following Answer based only on verified facts as of 2025, generate randomly, No extra text just return the JSON array:
				Subject: %s
Topic: %s
Format: %s
Difficulty: %s
Number of Questions: %d

Rules:
1. If Format = MCQ:
   - Return a JSON array of objects exactly like:
     [
       {
         "question": "...",
         "options": ["...","...","...","..."],
         "answer": "..." 
       }
     ]
   - Provide exactly 4 options for each question, only one correct.
2. If Format = True/False:
   - Return a JSON array of objects exactly like:
     [
       {
         "question": "...",
         "answer": "True"
       }
     ]
   - Use string "True" or "False" for answers.
3. If Format = Question:
   - Return a JSON array of objects exactly like:
     [
       {
         "question": "...",
         "answer": "..."
       }
     ]
4. All output MUST be valid JSON and contain only the JSON array (no commentary or code fences).
5. Ensure question difficulty matches the requested level.
6. Keep content relevant to the Subject and Topic when provided.
7. No extra text just the JSON array.
`, subject, topic, format, GenrateQuizDTO.Difficulty, 10)
}

func (s *quizImplService) GenerateQuiz(GenrateQuizDTO *GenrateQuizDTO) (interface{}, error) {

	llm, err := ollama.New(ollama.WithModel("mistral"))
	if err != nil {
		log.Printf("Ollama init error: %v", err)
		return nil, err
	}
	format, prompt := buildQuizPrompt(GenrateQuizDTO)
	ctx := context.Background()
	resp, err := llm.Call(ctx, prompt)
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
