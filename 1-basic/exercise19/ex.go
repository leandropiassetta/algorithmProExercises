package exercise19

import (
	"fmt"
)

func RunExercise19() {
	var start, end int

	fmt.Println("Please choose the range of numbers for which you'd like to generate the multiplication table.")
	fmt.Print("Enter the starting number: ")
	fmt.Scanln(&start)

	fmt.Print("Enter the ending number: ")
	fmt.Scanln(&end)

	fmt.Printf("  |")
	for i := start; i <= end; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println("\n-------------------------------------------")

	for i := start; i <= end; i++ {
		fmt.Printf("%2d|", i)
		for j := start; j <= end; j++ {
			result := i * j
			fmt.Printf("%4d", result)
		}
		fmt.Println()
	}
}
