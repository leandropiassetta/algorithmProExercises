package exercise13

import "fmt"

func RunExercise13() {
	var year int

	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	// pass 1: Determine if the number is divisible for 4

	isLeapYear := year%4 == 0 && year%100 != 0 || year%100 == 0 && year%400 == 0

	if !isLeapYear {
		fmt.Println("not is a leap year")
		return
	} else {
		fmt.Println("is a leap year")
		return
	}
}
