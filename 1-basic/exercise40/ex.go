package exercise40

import (
	"fmt"
)

func RunExercise40() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	specifierChar := "F"
	endsCharacter := str[len(str)-1]

	if string(endsCharacter) == specifierChar {
		fmt.Printf("This string: %s ends with %s\n", str, specifierChar)
	} else {
		fmt.Printf("This string: %s does not ends with %s\n", str, specifierChar)
	}
}
