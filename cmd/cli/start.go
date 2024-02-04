package cli

import (
	"fmt"
	"os"

	"github.com/Allyedge/allyignore/gitignore"
	"github.com/Allyedge/allyignore/util"
)

func Start(noKeep bool) {
	files := gitignore.ReadDirectory(".")

	gitignoreIndex := util.Find(files, ".gitignore")

	if gitignoreIndex+1 > len(files) {
		fmt.Println("No .gitignore found, exiting!")
		os.Exit(0)
	}

	gitignoreFile := files[gitignoreIndex]

	lines := gitignore.ReadLines(gitignoreFile)

	uselessLines := gitignore.AddUselessLines(lines, noKeep)

	lines = gitignore.AddCleanLines(lines, uselessLines)

	content, err := os.Create(gitignoreFile)
	if err != nil {
		fmt.Println("Error creating .gitignore file!")
		os.Exit(0)
	}

	for _, line := range lines {
		_, err := fmt.Fprintln(content, line)
		if err != nil {
			fmt.Println("Error writing to .gitignore file!")
			os.Exit(0)
		}
	}

	if len(uselessLines) == 0 {
		fmt.Println("No useless lines found in .gitignore!")
		os.Exit(0)
	}

	fmt.Println("Deleted these files from .gitignore:")

	for _, line := range uselessLines {
		fmt.Println(line)
	}

	content.Close()
}
