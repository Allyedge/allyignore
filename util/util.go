package util

func Find(array []string, value string) int {
	for index, element := range array {
		if value == element {
			return index
		}
	}

	return len(array)
}
