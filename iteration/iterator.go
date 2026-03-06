package iteration

import "strings"

func Iterator(word string, multiplier int) string {
	var result strings.Builder
	for i := 0; i < multiplier; i++ {
		result.WriteString(word)
	}
	return result.String()
}
