package main

import (
	"fmt"
)

// exercise 2: Calculate the sum of two numbers.

func main() {
	var numberA int
	var numberB int

	fmt.Print("Enter number A: ")

	fmt.Scan(&numberA)

	fmt.Print("Enter number B: ")

	fmt.Scan(&numberB)

	fmt.Printf("The sum of %d and %d is %d \n", numberA, numberB, numberA+numberB)
}
