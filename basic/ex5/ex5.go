package main

import "fmt"

// exercise 5: Check if a number is positive, negative, or zero.

func main() {
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
