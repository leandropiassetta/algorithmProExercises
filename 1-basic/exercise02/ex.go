package exercise02

import "fmt"

func RunExercise02() {
	var numberA int
	var numberB int

	fmt.Print("Enter number A: ")

	fmt.Scan(&numberA)

	fmt.Print("Enter number B: ")

	fmt.Scan(&numberB)

	fmt.Printf("The sum of %d and %d is %d \n", numberA, numberB, numberA+numberB)
}
