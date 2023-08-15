package main

import (
	"fmt"

	"github.com/leandropiassetta/algorithmProExercises/basic/ex1"
	"github.com/leandropiassetta/algorithmProExercises/basic/ex2"
	"github.com/leandropiassetta/algorithmProExercises/basic/ex3"
	"github.com/leandropiassetta/algorithmProExercises/basic/ex4"
	"github.com/leandropiassetta/algorithmProExercises/basic/ex5"
	"github.com/leandropiassetta/algorithmProExercises/basic/ex6"
)

func main() {
	fmt.Println("Basic Exercises Section")

	fmt.Println("exercise 1: Print " + "Hello, World!" + "to the console")
	ex1.RunExercise1()

	fmt.Println("exercise 2: Calculate the sum of two numbers.")
	ex2.RunExercise2()

	fmt.Println("exercise 3: Calculate the average of three numbers.")
	ex3.RunExercise3()

	fmt.Println("exercise 4: Check if a number is even or odd.")
	ex4.RunExercise4()

	fmt.Println("exercise 5: Check if a number is positive, negative, or zero.")
	ex5.RunExercise5()

	fmt.Println("exercise 6: Convert Celsius to Fahrenheit.")
	ex6.RunExercise6()
}
