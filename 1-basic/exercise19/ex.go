package exercise19

import (
	"fmt"
)

func RunExercise19() {
	// var start, end int

	// fmt.Println("Please choose the range of numbers for which you'd like to generate the multiplication table.")
	// fmt.Print("Enter the starting number: ")
	// fmt.Scanln(&start)

	// fmt.Print("Enter the ending number: ")
	// fmt.Scanln(&end)

	var header string

	start := 1
	end := 10

	for i := start; i < end+1; i++ {
		header += fmt.Sprintf("  |%v", i)
	}

	fmt.Println(header)
	fmt.Println("------------------------------------------")
	var line string

	for i := start; i < end+1; i++ {
		line += fmt.Sprintf("\n%v|", i)
		for j := start; j < end+1; j++ {
			line += fmt.Sprintf("  %v", j*i)
		}
	}

	fmt.Println(line)
}
