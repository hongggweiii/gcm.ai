package ui

import (
	"fmt"

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
		return userChoices, fmt.Errorf("Error running CLI form %v\n", err)
	}

	return userChoices, nil
}
