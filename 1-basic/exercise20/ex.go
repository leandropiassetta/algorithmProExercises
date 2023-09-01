package exercise20

import (
	"errors"
	"fmt"
	"math"
)

func foundRoots(a, b, c float64) (float64, float64, error) {
	d := delta(a, b, c)
	if d < 0 {
		return 0, 0, errors.New("complex roots")
	}
	rootA := (-b + math.Sqrt(d)) / (2 * a)
	rootB := (-b - math.Sqrt(d)) / (2 * a)

	return rootA, rootB, nil
}

func delta(a, b, c float64) float64 {
	delta := (b * b) - 4*a*c

	return delta
}

func RunExercise20() {
	// quadratic equation
	// f(x) = x2 -5x + 6

	rootA, rootB, err := foundRoots(1, -5, 6)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rootA, rootB)
	}
}
