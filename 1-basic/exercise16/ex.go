package exercise16

import "fmt"

func RunExercise16() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scan(&str)

	var reverseString string
	for i := len(str) - 1; i > -1; i-- {
		reverseString += string(str[i])
	}

	fmt.Printf("The reverse string is %s\n", reverseString)
}
