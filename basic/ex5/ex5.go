package ex5

import "fmt"

func RunExercise5() {
	var number float64

	fmt.Print("Digit a number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Print("This number is greater than zero\n")
	} else if number < 0 {
		fmt.Print("This number is less than zero\n")
	} else {
		fmt.Print("This is equal to zero\n")
	}
}
