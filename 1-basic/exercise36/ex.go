package exercise36

import (
	"fmt"
	"strings"
)

func RunExercise36() {
	var str string
	var vowelsCount int

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	str = strings.ToLower(str)

	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsCount++
		}
	}

	fmt.Printf("The amount of vowels in this string: %s is %d\n", str, vowelsCount)
}
