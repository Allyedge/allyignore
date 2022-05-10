package internal

import (
	"fmt"
	"os"
)

func Start(noKeep bool) {
	var files = ReadDirectory(".")

	var gitignoreIndex = Find(files, ".gitignore")

	if gitignoreIndex+1 > len(files) {
		fmt.Println("No .gitignore found, exiting!")
		os.Exit(0)
	}

	var gitignoreFile = files[gitignoreIndex]

	var lines = ReadLines(gitignoreFile)

	var uselessLines = AddUselessLines(lines, noKeep)

	lines = AddCleanLines(lines, uselessLines)

	content, err := os.Create(gitignoreFile)

	CheckError(err, false)

	for _, line := range lines {
		_, err := fmt.Fprintln(content, line)
		CheckError(err, false)
	}

	content.Close()
}
