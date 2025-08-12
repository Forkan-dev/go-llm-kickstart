// ==========================================
// LLM Service Package (llm/service.go)
// ==========================================
package llm

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"
)

// LLMService interface defines core LLM operations

// GenerationRequest represents a generic LLM generation request
type GenerationRequest struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	TopK             int      `json:"top_k"`
	RepeatPenalty    float64  `json:"repeat_penalty"`
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	MaxTokens        int      `json:"max_tokens"`
	Stop             []string `json:"stop"`
	Stream           bool     `json:"stream"`
	Format           string   `json:"format"`
}

// RandomizationConfig holds parameters for different randomization strategies
type RandomizationConfig struct {
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	TopK             int      `json:"top_k"`
	RepeatPenalty    float64  `json:"repeat_penalty"`
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	MaxTokens        int      `json:"max_tokens"`
	Stop             []string `json:"stop"`
	CustomSeed       string   `json:"custom_seed,omitempty"`
}

// LLMResponse represents the response from LLM API
type LLMResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
	Error    string `json:"error,omitempty"`
}

// llmServiceImpl implements the LLMService interface
type llmServiceImpl struct {
	httpClient *http.Client
	model      string
}

// NewLLMService creates a new LLM service instance
func NewLLMService(model string) LLMService {
	if model == "" {
		model = "llama3.2:latest"
	}

	return &llmServiceImpl{
		httpClient: &http.Client{Timeout: 60 * time.Second},
		model:      model,
	}
}

// GenerateRandomSeed creates a random seed for reproducible randomization
func (s *llmServiceImpl) generateRandomSeed() string {
	timestamp := time.Now().Unix()
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(9999))
	return fmt.Sprintf("%d%d", timestamp, randomNum.Int64())
}

// GenerateResponse sends a generic request to the LLM API
func (s *llmServiceImpl) GenerateResponse(baseURL string, request GenerationRequest) (*LLMResponse, error) {
	if request.Model == "" {
		request.Model = s.model
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	url := fmt.Sprintf("%s/api/generate", baseURL)
	resp, err := s.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var llmResponse LLMResponse
	if err := json.Unmarshal(body, &llmResponse); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if llmResponse.Error != "" {
		return nil, fmt.Errorf("LLM API error: %s", llmResponse.Error)
	}

	return &llmResponse, nil
}

// HighImpactRandomize implements high-impact randomization strategy
func (s *llmServiceImpl) HighImpactRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        prompt,
		Temperature:   0.75,
		TopP:          0.95,
		TopK:          70,
		RepeatPenalty: 1.3,
		MaxTokens:     config.MaxTokens,
		Stop:          []string{"Human:", "\n\nHuman:", "```", "---"},
		Stream:        false,
	}

	// Add presence and frequency penalties for high impact
	presencePenalty := 0.6
	frequencyPenalty := 0.8
	request.PresencePenalty = &presencePenalty
	request.FrequencyPenalty = &frequencyPenalty

	return s.GenerateResponse(baseURL, request)
}

// ExtremeRandomize implements extreme randomization strategy
func (s *llmServiceImpl) ExtremeRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        prompt,
		Temperature:   0.85,
		TopP:          0.98,
		TopK:          100,
		RepeatPenalty: 1.4,
		MaxTokens:     config.MaxTokens,
		Stream:        false,
	}

	return s.GenerateResponse(baseURL, request)
}

// SeedBasedRandomize implements seed-based randomization
func (s *llmServiceImpl) SeedBasedRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	seed := config.CustomSeed
	if seed == "" {
		seed = s.generateRandomSeed()
	}

	enhancedPrompt := fmt.Sprintf("Session ID: %s. %s", seed, prompt)

	request := GenerationRequest{
		Model:         s.model,
		Prompt:        enhancedPrompt,
		Temperature:   0.8,
		TopP:          0.95,
		TopK:          80,
		RepeatPenalty: 1.35,
		MaxTokens:     config.MaxTokens,
		Stream:        false,
	}

	return s.GenerateResponse(baseURL, request)
}

// ContextSwitchingRandomize implements context-switching strategy
func (s *llmServiceImpl) ContextSwitchingRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        prompt,
		Temperature:   0.82,
		TopP:          0.96,
		TopK:          85,
		RepeatPenalty: 1.4,
		MaxTokens:     config.MaxTokens,
		Stream:        false,
	}

	return s.GenerateResponse(baseURL, request)
}

// MultiAngleRandomize generates from different perspectives
func (s *llmServiceImpl) MultiAngleRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        prompt,
		Temperature:   0.78,
		TopP:          0.93,
		TopK:          75,
		RepeatPenalty: 1.32,
		MaxTokens:     config.MaxTokens,
		Stream:        false,
	}

	return s.GenerateResponse(baseURL, request)
}

// UltimateRandomize implements the most random strategy
func (s *llmServiceImpl) UltimateRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        prompt,
		Temperature:   0.9,
		TopP:          0.98,
		TopK:          100,
		RepeatPenalty: 1.5,
		MaxTokens:     config.MaxTokens,
		Stream:        false,
	}

	return s.GenerateResponse(baseURL, request)
}

func (s *llmServiceImpl) ValidateJsonStringwithLLm(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error) {
	request := GenerationRequest{
		Model:         s.model,
		Prompt:        "strictly Validate json and currect this json string make sure start with { end with }" + prompt,
		Temperature:   0.75,
		TopP:          0.95,
		TopK:          70,
		RepeatPenalty: 1.3,
		MaxTokens:     config.MaxTokens,
		Stop:          []string{"Human:", "\n\nHuman:", "```", "---"},
		Stream:        false,
		Format:        "json",
	}

	return s.GenerateResponse(baseURL, request)
}
