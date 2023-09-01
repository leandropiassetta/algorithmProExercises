package exercise21

import (
	"fmt"
	"strconv"
)

func RunExercise21() {
	var number string

	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)

	n, _ := strconv.Atoi(number)

	sum := 1

	for i := 1; i < n; i++ {

		nextNumber := i + 1

		sum += nextNumber
	}

	fmt.Print("The sum of natural numbers up to ", n, " is ", sum, "!\n")
}
