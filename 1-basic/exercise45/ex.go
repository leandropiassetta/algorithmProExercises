package exercise45

import (
	"fmt"
)

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

func RunExercise45() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	fmt.Printf("The factorial of %d is %d\n", number, factorial(number))
}
