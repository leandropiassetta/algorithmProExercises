package exercise47

import (
	"fmt"
	"strconv"
)

func RunExercise47() {
	var number int

	fmt.Print("Enter the number : ")
	fmt.Scan(&number)

	var fibonacciSequence string
	previous := 0
	current := 1

	for i := 1; i <= number; i++ {
		fibonacciSequence += strconv.Itoa(current) + ", "
		next := current + previous
		previous = current
		current = next
	}

	fmt.Printf("The Fibonacci sequence for this term is: %s\n", fibonacciSequence[:len(fibonacciSequence)-2])

	fmt.Printf("the next fibonacci number after of %d is %d\n", number, current)
}
