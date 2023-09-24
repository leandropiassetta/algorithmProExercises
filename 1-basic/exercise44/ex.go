package exercise44

import (
	"fmt"
)

func isPowerOfTwo(number int) bool {
	if number <= 0 {
		return false
	}

	for number > 1 {
		if number%2 != 0 {
			return false
		}
		number = number / 2
	}

	return true
}

func RunExercise44() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	if isPowerOfTwo(number) {
		fmt.Printf("The number %d is a power of 2\n", number)
	} else {
		fmt.Printf("The number %d is not a power of 2\n", number)
	}
}
