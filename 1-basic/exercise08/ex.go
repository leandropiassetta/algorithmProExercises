package exercise08

import "fmt"

func RunExercise08() {
	var number int

	fmt.Print("Enter the 1º number: ")
	fmt.Scan(&number)
	smallNumber := number

	for i := 2; i <= 3; i++ {
		fmt.Printf("Enter the %dº number: ", i)
		fmt.Scan(&number)

		if number < smallNumber {
			smallNumber = number
		}
	}
	fmt.Printf("The smallest number is: %d\n", smallNumber)
}
