package main

import "fmt"

// exercise 4: Check if a number is even or odd.

func main() {
	var number int

	fmt.Print("Digit a number: ")
	fmt.Scan(&number)

	even := number%2 == 0

	if even {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}
