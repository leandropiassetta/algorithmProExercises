package exercise50

import (
	"fmt"
)

func RunExercise50() {
	var number int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading the number.")
		return
	}
}
