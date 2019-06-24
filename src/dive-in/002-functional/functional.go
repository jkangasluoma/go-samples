package main

import (
	"fmt"
	"strings"
)

func main() {
	thStripper := stripper(youGo, "th")
	fmt.Println(thStripper)
}

// Higher order functions:
var stripper = func(f func() []string, strippedCharacters string) string {
	result := strings.Join(f(), " ")

	for i := 0; i < len(strippedCharacters); i++ {
		result = strings.ReplaceAll(result, string(strippedCharacters[i]), "")
	}
	return result
}

func youGo() []string {
	// Immutables:
	const imperative = "Go"
	const subject = "Youth!"
	return []string{imperative, subject}
}
