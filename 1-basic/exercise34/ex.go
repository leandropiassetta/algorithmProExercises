package exercise34

import "fmt"

func ToUpper(input string) string {
	result := ""

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		result += string(char)
	}

	return result
}

func RunExercise34() {
	str := "Hello, World!"
	upperStr := ToUpper(str)
	fmt.Println(upperStr)
}
