package gitignore

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Allyedge/allyignore/util"
)

func specialUselessLine(line string) bool {
	return strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0
}

func ReadDirectory(readPath string) []string {
	var directoryFiles []string

	files, err := ioutil.ReadDir(readPath)
	if err != nil {
		fmt.Println("Error reading directory!")
		os.Exit(0)
	}

	for _, file := range files {
		directoryFiles = append(directoryFiles, file.Name())
	}

	return directoryFiles
}

func ReadLines(gitignore string) []string {
	var lines []string

	content, err := os.Open(gitignore)

	if err != nil {
		fmt.Println("Error opening .gitignore file!")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err != nil {
		fmt.Println("Error opening .gitignore file!")
		os.Exit(0)
	}

	content.Close()

	return lines
}

func AddUselessLines(lines []string, noKeep bool) []string {
	var uselessLines []string

	for _, line := range lines {
		if _, err := os.Stat(line); os.IsNotExist(err) {
			if noKeep {
				uselessLines = append(uselessLines, line)
			} else {
				if !specialUselessLine(line) {
					uselessLines = append(uselessLines, line)
				}
			}

			continue
		}
	}

	return uselessLines
}

func AddCleanLines(lines []string, uselessLines []string) []string {
	for _, uselessLine := range uselessLines {
		var index = util.Find(lines, uselessLine)

		if index+1 > len(lines) {
			continue
		}

		lines = append(lines[:index], lines[index+1:]...)
	}

	if len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}

	return lines
}
