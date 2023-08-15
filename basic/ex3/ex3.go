package ex3

import "fmt"

func RunExercise3() {
	var numberA int
	var numberB int
	var numberC int

	fmt.Println("Enter the first number: ")
	fmt.Scan(&numberA)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&numberB)

	fmt.Println("Enter the third number: ")
	fmt.Scan(&numberC)

	sumNumbers := numberA + numberB + numberC
	average := sumNumbers / 3

	fmt.Printf("The average of this numbers: %v, %v, %v is: %v", numberA, numberB, numberC, average)
}
