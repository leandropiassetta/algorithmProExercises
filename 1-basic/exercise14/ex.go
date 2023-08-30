package exercise14

import "fmt"

func RunExercise14() {
	var str string

	fmt.Println("Enter a string or leave it blank")
	fmt.Scanln(&str)

	if str == "" {
		fmt.Println("This string is blank")
	} else {
		fmt.Printf("The chosen string is: %s", str)
	}
}
