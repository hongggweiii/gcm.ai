package main

import (
	"fmt"
	"os"

	"github.com/hongggweiii/gcm.ai/internal/git"
	"github.com/hongggweiii/gcm.ai/internal/ui"
)

func main() {
	choices, err := ui.FormConfig()
	if err != nil {
		fmt.Println("User cancelled the form")
		return
	}

	gitTool := git.NewLocalGit()
	diff, err := gitTool.GetStagedDiff()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch staged diff: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n--- Verification ---\n")
	fmt.Printf("Style: %v (True means Conventional)\n", choices.IsConventional)
	fmt.Printf("Multi-line: %v\n", choices.IsSingleLine)
	fmt.Printf("Changes found: %s", diff)
}
