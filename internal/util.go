package internal

import (
	"fmt"
)

func CheckError(err error, shouldContinue bool) {
	if err != nil {
		if shouldContinue {
			return
		} else {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func Find(array []string, value string) int {
	for index, element := range array {
		if value == element {
			return index
		}
	}

	return len(array)
}
