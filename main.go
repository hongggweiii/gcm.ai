package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hongggweiii/gcm.ai/internal/git"
	"github.com/hongggweiii/gcm.ai/internal/llm"
	"github.com/hongggweiii/gcm.ai/internal/ui"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	choices, err := ui.ConfigForm()
	if err != nil {
		fmt.Println("User cancelled the form")
		return
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	client := llm.NewGeminiClient(os.Getenv("GEMINI_API_KEY"), "gemini-2.5-flash")

	gitTool := git.NewLocalGit()
	diff, err := gitTool.GetStagedDiff()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch staged diff: %v\n", err)
		os.Exit(1)
	}

	suggestions, err := client.GenerateSuggestions(ctx, diff, true, true)

	selected, err := ui.PickMessageForm(suggestions)
	if err != nil {
		fmt.Println("User cancelled the form")
		return
	}

	isCommit, err := ui.CommitMessageForm()
	if err != nil {
		fmt.Println("User cancelled the form")
		return
	}

	fmt.Printf("isConventional: %t\n", choices.IsConventional)
	fmt.Printf("isSingleLine: %t\n", choices.IsSingleLine)
	fmt.Printf("Suggestions given: %s\n", suggestions)
	fmt.Printf("Selected commit message: %s\n", selected)

	if isCommit {
		isSuccess, err := gitTool.CommitMessage(selected)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to commit changes: %v\n", err)
			os.Exit(1)
		}

		if isSuccess {
			fmt.Println("Changes are successfully committed")
		}
	}
}
