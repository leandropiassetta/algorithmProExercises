package exercise33

import (
	"fmt"
)

func containsSubstring(mainString, substring string) bool {
	mainStringLength := len(mainString)
	substringLength := len(substring)

	for i := 0; i <= mainStringLength-substringLength; i++ {
		match := true

		for j := 0; j < substringLength; j++ {
			if mainString[i+j] != substring[j] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

func RunExercise33() {
	var str string
	var substr string

	fmt.Println("Enter a string:")
	fmt.Scan(&str)

	fmt.Println("Enter a substring:")
	fmt.Scan(&substr)

	if containsSubstring(str, substr) {
		fmt.Printf("The string %v contains the substring %v\n", str, substr)
	} else {
		fmt.Printf("The string %v does not contain the substring %v\n", str, substr)
	}
}
