package exercise11

import (
	"fmt"
	"time"
)

// Function to calculate factorial without using recursion
func calculateFactorialIterative(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Function to calculate factorial using recursion
func calculateFactorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * calculateFactorialRecursive(n-1)
}

func RunExercise11() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	time.Sleep(1 * time.Second)

	// Calculate and display factorial using the recursive function
	fmt.Printf("Factorial of %d (recursive) is %d\n", num, calculateFactorialRecursive(num))
	// Calculate and display factorial using the non-recursive function
	fmt.Printf("Factorial of %d (non-recursive) is %d\n", num, calculateFactorialIterative(num))
}
