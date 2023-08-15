package basic

import "fmt"

func runExercise1() {
	fmt.Println("Hello, World!")
}

func runExercise2() {
	var numberA int
	var numberB int

	fmt.Print("Enter number A: ")

	fmt.Scan(&numberA)

	fmt.Print("Enter number B: ")

	fmt.Scan(&numberB)

	fmt.Printf("The sum of %d and %d is %d \n", numberA, numberB, numberA+numberB)
}

func runExercise3() {
	var numberA int
	var numberB int
	var numberC int

	fmt.Println("Enter the first number: ")
	fmt.Scan(&numberA)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&numberB)

	fmt.Println("Enter the third number: ")
	fmt.Scan(&numberC)

	sumNumbers := numberA + numberB + numberC
	average := sumNumbers / 3

	fmt.Printf("The average of this numbers: %v, %v, %v is: %v", numberA, numberB, numberC, average)
}

func runExercise4() {
	var number int

	fmt.Print("Digit a number: ")
	fmt.Scan(&number)

	even := number%2 == 0

	if even {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
}

func runExercise5() {
	var number float64

	fmt.Print("Digit a number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Print("This number is greater than zero\n")
	} else if number < 0 {
		fmt.Print("This number is less than zero\n")
	} else {
		fmt.Print("This is equal to zero\n")
	}
}

func convertTemperature(celsius float64) float64 {
	fahrenheit := (celsius * 1.8) + 32

	return fahrenheit
}

func runExercise6() {
	var celsiusTemperature float64

	fmt.Print("Enter the current celsius temperature of your city? ")
	fmt.Scan(&celsiusTemperature)

	fahrenheitTemperature := convertTemperature(celsiusTemperature)

	fmt.Printf("The current fahrenheit temperature of your city is: %.1f\n", fahrenheitTemperature)
}

func runExercise7() {
	var numberA int
	var numberB int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&numberA)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&numberB)

	fmt.Printf("We have these two numbers: [%d, %d]\n", numberA, numberB)

	if numberA > numberB {
		fmt.Printf("The larger number of these two numbers is: %d\n", numberA)
	} else {
		fmt.Printf("The larger number of these two numbers is: %d\n", numberB)
	}
}

func runExercise8() {
	var number int

	fmt.Print("Enter the 1º number: ")
	fmt.Scan(&number)
	smallNumber := number

	for i := 2; i <= 3; i++ {
		fmt.Printf("Enter the %dº number: ", i)
		fmt.Scan(&number)

		if number < smallNumber {
			smallNumber = number
		}
	}
	fmt.Printf("The smallest number is: %d\n", smallNumber)
}

func RunExercisesBasics() {
	fmt.Println("exercise 1: Print " + "Hello, World!" + "to the console")
	runExercise1()

	fmt.Println("exercise 2: Calculate the sum of two numbers.")
	runExercise2()

	fmt.Println("exercise 3: Calculate the average of three numbers.")
	runExercise3()

	fmt.Println("exercise 4: Check if a number is even or odd.")
	runExercise4()

	fmt.Println("exercise 5: Check if a number is positive, negative, or zero.")
	runExercise5()

	fmt.Println("exercise 6: Convert Celsius to Fahrenheit.")
	runExercise6()

	fmt.Println("exercise 7: Find the larger of two numbers.")
	runExercise7()

	fmt.Println("exercise 8: Find the smaller of three numbers.")
	runExercise8()
}
