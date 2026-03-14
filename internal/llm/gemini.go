package llm

type GeminiClient struct {
	apiKey string
	model  string
}

func NewGeminiClient(apiKey string, model string) AIProvider {
	return &GeminiClient{
		apiKey: apiKey,
		model:  model,
	}
}

func (client *GeminiClient) GenerateSuggestions(diff string) ([]string, error) {
	return
}
