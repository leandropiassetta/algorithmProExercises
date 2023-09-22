package exercise35

import "fmt"

func ToLower(input string) string {
	result := ""

	for _, char := range input {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		result += string(char)
	}

	return result
}

func RunExercise35() {
	str := "HELLO, WORLD!"
	upperStr := ToLower(str)
	fmt.Println(upperStr)
}
