package llm

import (
	"context"
	"encoding/json"
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

	response, err := client.Models.GenerateContent(
		ctx,
		c.model,
		genai.Text(diff),
		&reqConfig,
	)

	if err != nil {
		fmt.Printf("Gemini unable to generate content: %v", err)
	}

	var suggestions CommitSuggestions

	// Extract JSON string from GenerateContentResponse wrapper
	rawText := response.Text()
	if err := json.Unmarshal([]byte(rawText), &suggestions); err != nil {
		return nil, fmt.Errorf("Failed to parse JSON: %w", err)
	}

	return suggestions.Suggestions, err
}
