package exercise12

import (
	"fmt"
	"time"
)

func RunExercise12() {
	var number int

	fmt.Print("Enter a number: \n")
	time.Sleep(1 * time.Second)

	fmt.Scan(&number)

	// Create a boolean slice of size 'number + 1' and fill with 'true'
	prime := make([]bool, number+1)
	for i := 2; i <= number; i++ {
		prime[i] = true
	}

	// Step 1: Exclude all even numbers except 2
	for i := 4; i <= number; i += 2 {
		prime[i] = false
	}

	// Step 2: Exclude multiples of 3 except 3
	for i := 6; i <= number; i += 3 {
		prime[i] = false
	}

	// Step 3: Exclude numbers ending in 0 or 5 except 5
	for i := 10; i <= number; i += 5 {
		prime[i] = false
	}
	for i := 15; i <= number; i += 10 {
		prime[i] = false
	}

	// Step 4: Exclude multiples of 7 except 7
	for i := 14; i <= number; i += 7 {
		prime[i] = false
	}

	// Print all prime numbers
	fmt.Println("Prime numbers up to", number, ":")
	for i := 2; i <= number; i++ {
		if prime[i] {
			fmt.Printf("%d ", i)
		}
	}
}
