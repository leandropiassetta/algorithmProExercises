package exercise32

import (
	"fmt"
	"math"
)

func RunExercise32() {
	var radius float64

	fmt.Println("Enter the radius of the circle:")
	fmt.Scan(&radius)

	perimeter := radius * (2 * math.Pi)

	fmt.Printf("The perimeter of a circle with radius %.2f is: %.2f\n", radius, perimeter)
}
