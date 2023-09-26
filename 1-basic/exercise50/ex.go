package exercise50

import (
	"fmt"
)

func RunExercise50() {
	var year int

	fmt.Print("Enter a year: ")
	_, err := fmt.Scan(&year)
	if err != nil {
		fmt.Println("Error reading the year.")
		return
	}

	firstWorldCup := 1930
	lastWorldCup := 2022

	isWorldCupYear := (year-firstWorldCup)%4 == 0 && year >= firstWorldCup && year <= lastWorldCup && (year != 1942 && year != 1946)

	if isWorldCupYear {
		fmt.Printf("this year: %d was a world cup year\n", year)
	} else {
		fmt.Printf("this year: %d was not a world cup year\n", year)
	}
}
