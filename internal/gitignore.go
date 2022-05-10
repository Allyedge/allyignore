package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func specialUselessLine(line string) bool {
	return strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0
}

func ReadDirectory(readPath string) []string {
	var directoryFiles []string

	files, err := ioutil.ReadDir(readPath)

	CheckError(err, false)

	for _, file := range files {
		directoryFiles = append(directoryFiles, file.Name())
	}

	return directoryFiles
}

func ReadLines(gitignore string) []string {
	var lines []string

	content, err := os.Open(gitignore)

	CheckError(err, false)

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	CheckError(scanner.Err(), false)

	content.Close()

	return lines
}

func AddUselessLines(lines []string, noKeep bool) []string {
	var uselessLines []string

	for _, line := range lines {
		if _, err := os.Stat(line); os.IsNotExist(err) {
			CheckError(err, true)

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
		var index = Find(lines, uselessLine)

		if index+1 > len(lines) {
			continue
		}

		lines = append(lines[:index], lines[index+1:]...)
	}

	fmt.Println(lines)

	if len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}

	return lines
}
