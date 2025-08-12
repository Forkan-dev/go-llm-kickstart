package llm

type LLMService interface {
	GenerateResponse(baseURL string, request GenerationRequest) (*LLMResponse, error)
	HighImpactRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	ExtremeRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	SeedBasedRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	ContextSwitchingRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	MultiAngleRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	UltimateRandomize(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
	ValidateJsonStringwithLLm(baseURL string, prompt string, config RandomizationConfig) (*LLMResponse, error)
}
