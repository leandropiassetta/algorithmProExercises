package exercise42

import (
	"fmt"
)

func isPerfectSquare(number int) (int, bool) {
	squareRoot := 0
	control := number

	if number < 0 {
		fmt.Println("The number is not a perfect square.")
		return 0, false
	}

	for i := 1; i < number; i += 2 {
		control = control - i

		if control == 0 {
			squareRoot++
			return squareRoot, true
		}
		squareRoot++
	}

	return 0, false
}

func RunExercise42() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	squareRoot, isPerfectSquare := isPerfectSquare(number)

	if isPerfectSquare {
		fmt.Printf("The number %d is a perfect square and its square root is %d\n", number, squareRoot)
	} else {
		fmt.Printf("The number %d is not a perfect square\n", number)
	}
}
