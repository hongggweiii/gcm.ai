package llm

import (
	"context"
)

type CommitSuggestions struct {
	Suggestions []string `json:"suggestions"` // Map field to key "suggestions"
}

type AIProvider interface {
	GenerateSuggestions(ctx context.Context, diff string, isConventional bool, isSingleLine bool) ([]string, error)
}
