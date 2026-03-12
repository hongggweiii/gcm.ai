package llm

type CommitSuggestions struct {
	Options []string `json:"options"` // Map field to key "options"
}

type APIProvider interface {
	GenerateSuggestions(diff string, isConventional bool, isSingleLine bool) ([]string, error)
}
