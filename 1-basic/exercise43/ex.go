package exercise43

import (
	"fmt"
	"math/rand"
)

func RunExercise43() {
	var firstNumber int
	fmt.Println("Enter the first number for the range: ")

	fmt.Scan(&firstNumber)

	var secondNumber int
	fmt.Println("Enter the second number for the range: ")

	fmt.Scan(&secondNumber)

	if firstNumber > secondNumber {
		fmt.Println("The first number must be less than the second number")
		return
	}

	randomNumber := rand.Intn(secondNumber-firstNumber+1) + firstNumber

	fmt.Printf("The random number between %d and %d is %d\n", firstNumber, secondNumber, randomNumber)
}
