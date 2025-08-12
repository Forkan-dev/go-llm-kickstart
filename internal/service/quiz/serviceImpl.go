// ==========================================
// Quiz Service Package (quiz/service.go)
// ==========================================
package quiz

import (
	"encoding/json"
	"fmt"
	"learning-companion/internal/service/llm"
	"log"
	"strings"
	"time"
)

// QuizService interface defines quiz-specific operations

// QuizConfig holds configuration for quiz generation
type QuizConfig struct {
	Subject       string   `json:"subject"`
	Difficulty    string   `json:"difficulty"`
	QuestionCount int      `json:"question_count"`
	QuestionTypes []string `json:"question_types"` // ["mcq", "true_false", "short_answer"]
	Topics        []string `json:"topics"`
	Strategy      string   `json:"strategy"` // "high_impact", "extreme", "seed_based", etc.
	MaxTokens     int      `json:"max_tokens"`
	CustomSeed    string   `json:"custom_seed,omitempty"`
}

// Question represents a single quiz question
type Question struct {
	Question      string   `json:"question"`
	Type          string   `json:"type"`
	Options       []string `json:"options,omitempty"`
	CorrectAnswer string   `json:"answer"`
	Explanation   string   `json:"explanation"`
	Difficulty    string   `json:"difficulty"`
	Topic         string   `json:"topic"`
}

// QuizResponse represents the structured quiz response
type QuizResponse struct {
	Questions []Question   `json:"questions"`
	Metadata  QuizMetadata `json:"metadata"`
}

// QuizMetadata holds additional information about the quiz
type QuizMetadata struct {
	GeneratedAt    time.Time `json:"generated_at"`
	Strategy       string    `json:"strategy"`
	Seed           string    `json:"seed,omitempty"`
	TotalQuestions int       `json:"total_questions"`
	Subject        string    `json:"subject"`
	Difficulty     string    `json:"difficulty"`
}

// quizServiceImpl implements the QuizService interface
type quizServiceImpl struct {
	llmService llm.LLMService
}

// NewQuizService creates a new quiz service instance
func NewQuizService(llmService llm.LLMService) Service {
	return &quizServiceImpl{
		llmService: llmService,
	}
}

// ValidateQuizConfig validates the quiz configuration
func (s *quizServiceImpl) ValidateQuizConfig(config QuizConfig) error {
	if config.QuestionCount <= 0 {
		return fmt.Errorf("question count must be greater than 0")
	}

	if config.Subject == "" {
		return fmt.Errorf("subject is required")
	}

	if len(config.QuestionTypes) == 0 {
		return fmt.Errorf("at least one question type is required")
	}

	validDifficulties := map[string]bool{
		"easy": true, "medium": true, "hard": true,
	}
	if !validDifficulties[config.Difficulty] {
		return fmt.Errorf("invalid difficulty level: %s", config.Difficulty)
	}

	validTypes := map[string]bool{
		"mcq": true, "true_false": true, "short_answer": true, "fill_blank": true, "true-false": true, "open-ended": true,
	}
	for _, qType := range config.QuestionTypes {
		if !validTypes[qType] {
			return fmt.Errorf("invalid question type: %s", qType)
		}
	}

	return nil
}

// BuildPrompt constructs the appropriate prompt for quiz generation
func (s *quizServiceImpl) BuildPrompt(config QuizConfig) (string, error) {
	var promptParts []string

	for _, qType := range config.QuestionTypes {
		switch qType {
		case "mcq":
			promptParts = append(promptParts, `{
  "question": "Your question here?",
  "type": "mcq",
  "options": ["Option A", "Option B", "Option C", "Option D"],
  "answer": "Option A",
  "explanation": "Clear explanation",
  "difficulty": "`+config.Difficulty+`",
  "topic": "specific topic"
}`)
		case "true_false", "true-false":
			promptParts = append(promptParts, `{
	make sure questions are can be answered by true or false, options are true and false and answer is true or false 
  "question": "Your true/false statement",
  "type": "true_false",
  "options": ["True", "False"],
  "answer": "True",
  "explanation": "Explanation of why this is true/false",
  "difficulty": "`+config.Difficulty+`",
  "topic": "specific topic"
}`)
		case "short_answer", "open-ended":
			promptParts = append(promptParts, `{
  "question": "Your open-ended question?",
  "type": "short_answer",
  "answer": "Expected answer",
  "explanation": "Additional context or explanation",
  "difficulty": "`+config.Difficulty+`",
  "topic": "specific topic"
}`)
		case "fill_blank":
			promptParts = append(promptParts, `{
  "question": "Complete this sentence: The capital of France is _____",
  "type": "fill_blank",
  "answer": "Paris",
  "explanation": "Paris is the capital and largest city of France",
  "difficulty": "`+config.Difficulty+`",
  "topic": "specific topic"
}`)
		}
	}

	basePrompt := fmt.Sprintf(`Subject: %s
Generate %d questions covering topics: %v
Difficulty: %s

Generate questions in this JSON format:
{"questions": [%s]}

Requirements:
- Each question must be unique and cover different concepts
- Use varied question structures
- Include real-world examples when possible
- Avoid typical textbook questions
- Make questions engaging and thought-provoking

Return only valid JSON.`,
		config.Subject,
		config.QuestionCount,
		config.Topics,
		config.Difficulty,
		promptParts[0])

	return basePrompt, nil
}

// ParseQuizResponse parses the LLM response into structured quiz data
func (s *quizServiceImpl) ParseQuizResponse(llmResponse string) (*QuizResponse, error) {
	var quizData struct {
		Questions []Question `json:"questions"`
	}
	log.Default().Println("llm response-------------->", llmResponse)
	start := strings.Index(llmResponse, "{")
	end := strings.LastIndex(llmResponse, "}")
	llmResponse = llmResponse[start : end+1]

	log.Default().Println("llm response-------------->", llmResponse)
	if err := json.Unmarshal([]byte(llmResponse), &quizData); err != nil {
		return nil, fmt.Errorf("error parsing quiz JSON: %w", err)
	}

	return &QuizResponse{
		Questions: quizData.Questions,
	}, nil
}

// GenerateQuiz is the main method that orchestrates quiz generation
func (s *quizServiceImpl) GenerateQuiz(baseURL string, config QuizConfig) (*QuizResponse, error) {
	// Validate configuration
	if err := s.ValidateQuizConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// Set defaults
	if config.MaxTokens == 0 {
		config.MaxTokens = 1200
	}
	if config.Strategy == "" {
		config.Strategy = "high_impact"
	}

	// Build prompt
	prompt, err := s.BuildPrompt(config)
	if err != nil {
		return nil, fmt.Errorf("error building prompt: %w", err)
	}

	// Create randomization config for LLM service
	randomizationConfig := llm.RandomizationConfig{
		MaxTokens:  config.MaxTokens,
		CustomSeed: config.CustomSeed,
	}

	// Call appropriate LLM strategy
	var response *llm.LLMResponse
	var strategy string = config.Strategy

	switch strategy {
	case "high_impact":
		response, err = s.llmService.HighImpactRandomize(baseURL, prompt, randomizationConfig)
	case "extreme":
		response, err = s.llmService.ExtremeRandomize(baseURL, prompt, randomizationConfig)
	case "seed_based":
		response, err = s.llmService.SeedBasedRandomize(baseURL, prompt, randomizationConfig)
	case "context_switching":
		response, err = s.llmService.ContextSwitchingRandomize(baseURL, prompt, randomizationConfig)
	case "multi_angle":
		response, err = s.llmService.MultiAngleRandomize(baseURL, prompt, randomizationConfig)
	case "ultimate":
		response, err = s.llmService.UltimateRandomize(baseURL, prompt, randomizationConfig)
	default:
		response, err = s.llmService.HighImpactRandomize(baseURL, prompt, randomizationConfig)
		strategy = "high_impact"
	}

	response, err = s.llmService.ValidateJsonStringwithLLm(baseURL, response.Response, randomizationConfig)
	
	if err != nil {
		return nil, fmt.Errorf("error generating quiz: %w", err)
	}

	// Parse response
	quizResponse, err := s.ParseQuizResponse(response.Response)
	if err != nil {
		return nil, fmt.Errorf("error parsing quiz response: %w", err)
	}

	// Add metadata
	quizResponse.Metadata = QuizMetadata{
		GeneratedAt:    time.Now(),
		Strategy:       strategy,
		Seed:           config.CustomSeed,
		TotalQuestions: len(quizResponse.Questions),
		Subject:        config.Subject,
		Difficulty:     config.Difficulty,
	}

	return quizResponse, nil
}
