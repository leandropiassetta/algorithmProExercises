package exercise38

import (
	"fmt"
)

func RunExercise38() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	count := 0

	for range str {
		count++
	}

	fmt.Printf("The length of this string: %s is %d\n", str, count)
}
