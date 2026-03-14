package llm

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiClient struct {
	apiKey       string
	model        string
	systemPrompt string
	config       genai.GenerateContentConfig
}

func NewGeminiClient(apiKey string, model string) AIProvider {
	basePrompt := SystemPrompt()

	return &GeminiClient{
		apiKey:       apiKey,
		model:        model,
		systemPrompt: basePrompt,
		config: genai.GenerateContentConfig{
			ResponseMIMEType: "application/json",
		},
	}
}

func (c *GeminiClient) GenerateSuggestions(ctx context.Context, diff string, isConv bool, isSingle bool) ([]string, error) {
	// Initialise new Gemini client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: c.apiKey,
	})
	if err != nil {
		fmt.Printf("Error initialising Gemini Client: %v", err)
	}

	// Set up prompt and send it to Gemini for suggestions
	finalPrompt := fmt.Sprintf(c.systemPrompt, isConv, isSingle)
	reqConfig := c.config
	reqConfig.SystemInstruction = genai.NewContentFromText(finalPrompt, "user")

	suggestions, err := client.Models.GenerateContent(
		ctx,
		c.model,
		genai.Text(diff),
		&reqConfig,
	)

	if err != nil {
		fmt.Printf("Gemini unable to generate content: %v", err)
	}
	fmt.Println(suggestions.Text())

	return []string{suggestions.Text()}, err
}
