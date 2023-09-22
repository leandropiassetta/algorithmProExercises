package exercise37

import (
	"fmt"
	"strings"
)

func RunExercise37() {
	var str string
	var consonantsCount int

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	str = strings.ToLower(str)
	vowels := "aeiou"

	for _, char := range str {
		if !strings.ContainsRune(vowels, char) {
			consonantsCount++
		}
	}

	fmt.Printf("The amount of consonants in this string: %s is %d\n", str, consonantsCount)
}
