package main

import (
	"fmt"
	"log"

	"github.com/hongggweiii/gcm.ai/internal/git"
)

func main() {
	gitTool := git.NewLocalGit()
	diff, err := gitTool.GetStagedDiff()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Changes found: %s", diff)
}
