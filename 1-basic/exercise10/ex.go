package exercise10

import (
	"fmt"
	"time"
)

func RunExercise10() {
	var a string
	var b string

	fmt.Println("Fill in the variables a and b with the values you want to swap")
	time.Sleep(1 * time.Second)

	fmt.Print("Fill the variable a:\n")
	fmt.Scan(&a)

	fmt.Printf("This is the value of a: %v\n", a)
	time.Sleep(1 * time.Second)

	fmt.Print("Fill the variable b:\n")
	fmt.Scan(&b)

	fmt.Printf("This is the value of b: %v\n", b)
	time.Sleep(1 * time.Second)

	a, b = b, a

	fmt.Printf("After swapping, the value of a is now: %v\n", a)
	fmt.Printf("After swapping, the value of b is now: %v\n", b)
}
