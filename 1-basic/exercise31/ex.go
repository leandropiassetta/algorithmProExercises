package exercise31

import "fmt"

func RunExercise31() {
	// perimeter of a rectangle is the sum of all sides

	var length int
	var width int

	fmt.Println("Enter the length of the rectangle:")
	fmt.Scan(&length)

	fmt.Println("Enter the width of the rectangle:")
	fmt.Scan(&width)

	perimeter := 2 * (length + width)

	fmt.Printf("The perimeter of the rectangle with length %v and width %v is: %v\n", length, width, perimeter)
}
