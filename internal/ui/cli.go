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

func FormConfig() (Config, error) {
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

func PickMessage(suggestions []string) (string, error) {
	var selected string

	// Set options to be a Option of AI-generated commit message
	options := []huh.Option[string]{
		huh.NewOption("Cancel", "Cancel"),
	}

	for _, msg := range suggestions {
		options = append(options, huh.NewOption(msg, msg))
	}

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
