package exercise39

import (
	"fmt"
)

func RunExercise39() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	specifierChar := "F"
	startCharacter := str[0]

	if string(startCharacter) == specifierChar {
		fmt.Printf("This string: %s starts with %s\n", str, specifierChar)
	} else {
		fmt.Printf("This string: %s does not start with %s\n", str, specifierChar)
	}
}
