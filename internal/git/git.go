package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type GitProvider interface {
	GetStagedDiff() (string, error)
	CommitMessage(message string) error
}

type localGit struct{}

// Wrapping localGit to be a GitProvider implementation
func NewLocalGit() GitProvider {
	return &localGit{}
}

func (g *localGit) GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	stdout, err := cmd.Output()
	diff := string(stdout)

	if err != nil {
		return "", fmt.Errorf("Error getting staged diff: %v\n", err)
	}

	if diff == "" {
		return "", errors.New("No changes staged. Run 'git add' first")
	}

	return diff, nil
}

func (g *localGit) CommitMessage(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)

	// Piping Git console outputs to user terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error committing message: %v\n", err)
	}

	return nil
}
