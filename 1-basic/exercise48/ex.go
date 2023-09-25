package exercise48

import (
	"fmt"
	"math"
	"strconv"
)

func isNarcissistic(number int) bool {
	strNumber := strconv.Itoa(number)
	length := len(strNumber)
	sumOfDigitPowers := 0.0

	for _, digitChar := range strNumber {
		digit, err := strconv.Atoi(string(digitChar))
		if err != nil {
			fmt.Println("Error converting string to int")
			return false
		}

		sumOfDigitPowers += math.Pow(float64(digit), float64(length))
	}

	return number == int(sumOfDigitPowers)
}

func RunExercise48() {
	var number int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading the number.")
		return
	}

	if isNarcissistic(number) {
		fmt.Printf("%d is a narcissistic number.\n", number)
	} else {
		fmt.Printf("%d is not a narcissistic number.\n", number)
	}
}
