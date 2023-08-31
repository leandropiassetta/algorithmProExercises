package exercise18

import (
	"fmt"
	"math"
)

// func RunExercise18() {
// 	var number float64

// 	fmt.Println("Enter a number: ")
// 	fmt.Scanln(&number)

// 	squareRoot := math.Sqrt(number)

// 	if number < 0 {
// 		fmt.Println("Cannot calculate the square root of a negative number.")
// 	} else {
// 		fmt.Printf("The square root of %.2f is %.2f\n", number, squareRoot)
// 	}
// }

func RunExercise18() {
	var number float64

	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	squareRoot := math.Sqrt(number)

	if number < 0 {
		fmt.Println("Cannot calculate the square root of a negative number.")
	} else {
		fmt.Printf("The square root of %.2f is %.2f\n", number, squareRoot)
	}
}
