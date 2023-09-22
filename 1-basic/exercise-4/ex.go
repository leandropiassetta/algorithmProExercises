package exercise4

import "fmt"

func RunExercise4() {
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
