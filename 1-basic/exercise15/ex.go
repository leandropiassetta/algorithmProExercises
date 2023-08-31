package exercise15

import "fmt"

func RunExercise15() {
	var strA, strB string

	fmt.Println("Enter the first string:")
	_, err := fmt.Scanf("%s\n", &strA)
	if err != nil {
		fmt.Println("Error: only one string was entered")
		return
	}

	fmt.Println("Enter the second string:")
	_, err = fmt.Scanf("%s\n", &strB)
	if err != nil {
		fmt.Println("Error: only one string was entered", err)
		return
	}

	concatenatedString := strA + strB

	fmt.Printf("The concatenated string is: %s\n", concatenatedString)
}
