package ex7

import "fmt"

func RunExercise7() {
	var numberA int
	var numberB int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&numberA)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&numberB)

	fmt.Printf("We have these two numbers: [%d, %d]\n", numberA, numberB)

	if numberA > numberB {
		fmt.Printf("The larger number of these two numbers is: %d\n", numberA)
	} else {
		fmt.Printf("The larger number of these two numbers is: %d\n", numberB)
	}
}
