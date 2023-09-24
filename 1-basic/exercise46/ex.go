package exercise46

import (
	"fmt"
)

func RunExercise46() {
	var base float64
	var exponent int

	fmt.Print("Enter the base: ")
	fmt.Scan(&base)

	fmt.Print("Enter the exponent: ")
	fmt.Scan(&exponent)

	result := 1.0
	for i := 0; i < exponent; i += 1 {
		result *= base
	}

	fmt.Printf("%2.f^%d = %2.f\n", base, exponent, result)
}
