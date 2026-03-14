package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

// Store user preferences
type Config struct {
	IsConventional bool
	IsSingleLine   bool
}

func ConfigForm() (Config, error) {
	// Setting default values so 'Yes' gets highlighted first
	userChoices := Config{
		IsConventional: true,
		IsSingleLine:   true,
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("Conventional Style?").
				Options(
					huh.NewOption("Yes", true),
					huh.NewOption("No", false),
				).
				Height(6).
				Value(&userChoices.IsConventional),
		),

		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("Single line?\nSingle: Concise, for small changes\nMultiple: Detailed, for large changes").
				Options(
					huh.NewOption("Yes", true),
					huh.NewOption("No", false),
				).
				Height(6).
				Value(&userChoices.IsSingleLine),
		),
	)

	err := form.Run()
	if err != nil {
		return userChoices, fmt.Errorf("Error displaying commit config form: %v\n", err)
	}

	return userChoices, nil
}

func PickMessageForm(suggestions []string) (string, error) {
	var selected string

	// Populate suggestions into form options
	var options []huh.Option[string]
	for _, msg := range suggestions {
		options = append(options, huh.NewOption(msg, msg))
	}
	options = append(options, huh.NewOption("Cancel", "Cancel"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select one generated commit message").
				Options(
					options...,
				).
				Height(6).
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		return selected, fmt.Errorf("Error showing commit message suggestions: %v\n", err)
	}

	if selected == "Cancel" {
		fmt.Fprintf(os.Stderr, "Selection cancelled: %v\n", err)
		os.Exit(1)
	}

	return selected, nil
}

func CommitMessageForm() (bool, error) {
	var isCommit bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[bool]().
				Title("Would you like to commit?").
				Options(
					huh.NewOption("Yes", true),
					huh.NewOption("No", false),
				).
				Height(6).
				Value(&isCommit),
		),
	)

	err := form.Run()
	if err != nil {
		return isCommit, fmt.Errorf("Error showing commit message suggestions: %v\n", err)
	}

	if !isCommit {
		fmt.Fprintf(os.Stderr, "Changes are not committed: %v\n", err)
		os.Exit(1)
	}

	return isCommit, nil
}
