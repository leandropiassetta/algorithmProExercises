package exercise22

import (
	"fmt"
	"strconv"
	"strings"
)

func RunExercise22() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	var fibonacciSequence string
	previous := 0
	current := 1

	for i := 1; i <= number; i++ {
		fibonacciSequence += strconv.Itoa(current) + " ,"
		next := current + previous
		previous = current
		current = next
	}

	fmt.Printf("The Fibonacci sequence for this term is: %s", strings.TrimSuffix(fibonacciSequence, " ,"))
}
