package exercise41

import (
	"fmt"
	"strconv"
)

func RunExercise41() {
	var number float64

	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	// stringNumber := fmt.Sprintf("%f", number)

	sumDigits := 0
	numberStr := strconv.FormatFloat(number, 'f', -1, 64)

	for _, char := range numberStr {

		if char == '.' || char == ',' || char == '-' {
			continue
		}

		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("error: when converting string to int")
			return
		}

		sumDigits += digit
	}

	fmt.Printf("The sum of the digits of %f is %d\n", number, sumDigits)
}
