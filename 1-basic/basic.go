package basic

import (
	"fmt"
	"math"
	"strings"
)

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

func checkObtuse(a, b, c float64) bool {
	angleA := calculateAngle(a, b, c)
	angleB := calculateAngle(b, a, c)
	angleC := calculateAngle(c, a, b)

	anglesGreater90 := 0
	if angleA > 90 {
		anglesGreater90++
	}
	if angleB > 90 {
		anglesGreater90++
	}
	if angleC > 90 {
		anglesGreater90++
	}

	return anglesGreater90 == 1
}

func calculateAngle(a, b, c float64) float64 {
	cosA := (b*b + c*c - a*a) / (2 * b * c)
	return math.Acos(cosA) * 180 / math.Pi
}

func runExercise9() {
	var sideLengthA float64
	var sideLengthB float64
	var sideLengthC float64
	var base float64
	var height float64

	fmt.Print("Enter the side A length of the equilateral triangle: ")
	fmt.Scan(&sideLengthA)

	fmt.Print("Enter the side B length of the equilateral triangle: ")
	fmt.Scan(&sideLengthB)

	fmt.Print("Enter the side C length of the equilateral triangle: ")
	fmt.Scan(&sideLengthC)

	options := `
Choose an option to calculate the area of a triangle:
a - Equilateral triangle
b - Scalene triangle
c - Isosceles triangle
d - Right triangle
e - Acute triangle
f - Obtuse triangle
`
	var option string
	fmt.Println(options)
	fmt.Scan(&option)

	isTriangleValid := sideLengthA+sideLengthB > sideLengthC && sideLengthA+sideLengthC > sideLengthB && sideLengthB+sideLengthC > sideLengthA

	if !isTriangleValid {
		fmt.Println("the chosen sides do not form a triangle, please choose other sides")
	}

	// equilateral triangle has all sides equal
	equilateralTriangleIsValid := sideLengthA == sideLengthB && sideLengthB == sideLengthC
	equilateralTriangle := option == strings.ToLower("a") && equilateralTriangleIsValid

	if !equilateralTriangle {
		fmt.Println("this triangle is not equilateral because all sides are not equal, please choose side lengths that are equal or choose another option")
	}

	// scalene triangle has all sides different
	scaleneTriangleIsValid := sideLengthA != sideLengthB && sideLengthB != sideLengthC && sideLengthA != sideLengthC
	scaleneTriangle := option == strings.ToLower("b") && scaleneTriangleIsValid

	if !scaleneTriangle {
		fmt.Println("this triangle is not scalene because all sides are not different, please choose side lengths that are different or choose another option")
	}

	// isosceles triangle has two sides equal and one different side  (the different side is the base) and two equal angles (the angles opposite the equal sides) and the sum of the squares of the lengths of the legs is equal to the square of the length of the hypotenuse (Pythagorean theorem)

	isoscelesTriangleValid := sideLengthA == sideLengthB && sideLengthB != sideLengthC && sideLengthA != sideLengthC || sideLengthA == sideLengthC && sideLengthC != sideLengthB && sideLengthA != sideLengthB || sideLengthB == sideLengthC && sideLengthC != sideLengthA && sideLengthB != sideLengthA
	isoscelesTriangle := option == strings.ToLower("c") && isoscelesTriangleValid

	if !isoscelesTriangle {
		fmt.Println("this triangle is not isosceles because two sides are not equal and one different side, please choose side that fulfill this condition or choose another option")
	}

	// right triangle has one angle equal to 90 degrees and the sum of the squares of the lengths of the legs is equal to the square of the length of the hypotenuse (Pythagorean theorem)

	rightTriangleIsValid := math.Pow(sideLengthC, 2) == math.Pow(sideLengthA, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthA, 2) == math.Pow(sideLengthC, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthB, 2) == math.Pow(sideLengthA, 2)+math.Pow(sideLengthC, 2)
	rightTriangle := option == strings.ToLower("d") && rightTriangleIsValid

	if !rightTriangle {
		fmt.Println("this triangle is not right because one side is not equal that sum of the squares of the lengths of the legs, please choose side that fulfill this condition or choose another option")

		// acute triangle has all angles less than 90 degrees and the sum of the squares of the lengths of the legs is greater than the square of the length of the hypotenuse (Pythagorean theorem)

		acuteTriangleIsValid := math.Pow(sideLengthC, 2) < math.Pow(sideLengthA, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthA, 2) < math.Pow(sideLengthC, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthB, 2) < math.Pow(sideLengthA, 2)+math.Pow(sideLengthC, 2)
		acuteTriangle := (option == "e") && acuteTriangleIsValid

		if !acuteTriangle {
			fmt.Println("this triangle is not acute because one side is not less than the sum of the squares of the lengths of the legs, please choose side that fulfill this condition or choose another option")
		}

		// obtuse triangle has one angle greater than 90 degrees and the sum of the squares of the lengths of the legs is less than the square of the length of the hypotenuse (Pythagorean theorem)

		obtuseTriangle := option == strings.ToLower("f") && checkObtuse(sideLengthA, sideLengthB, sideLengthC)

		if !obtuseTriangle {
			fmt.Println("this triangle is not obtuse because one angle is not greater than 90 degrees, please choose side that fulfill this condition or choose another option")
		}

		switch {

		case equilateralTriangle:

			height = (math.Sqrt(3) / 2) * sideLengthA
			base = sideLengthC
			area := (base * height) / 2

			fmt.Printf("the area of equilateral triangle is: %.2f", area)

		case scaleneTriangle:
			perimeter := sideLengthA + sideLengthB + sideLengthC
			semiPerimeter := perimeter / 2

			area := math.Sqrt(semiPerimeter * (semiPerimeter - sideLengthA) * (semiPerimeter - sideLengthB) * (semiPerimeter - sideLengthC))

			fmt.Printf("the area of scalene triangle is: %.2f", area)

		case isoscelesTriangle:

			if sideLengthC == sideLengthB {
				base = sideLengthA
				// theorem of pythagoras
				height := math.Sqrt((sideLengthC * sideLengthC) - (base / 2 * base / 2))

				area := (base * height) / 2

				fmt.Printf("the area of isosceles triangle is: %.2f", area)
			} else if sideLengthA == sideLengthC {
				base = sideLengthB
				// theorem of pythagoras
				height := math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))

				area := (base * height) / 2

				fmt.Printf("the area of isosceles triangle is: %.2f", area)
			} else {
				base = sideLengthC
				// theorem of pythagoras
				height := math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))

				area := (base * height) / 2

				fmt.Printf("the area of isosceles triangle is: %.2f", area)
			}

		case rightTriangle:
			if sideLengthA < sideLengthB && sideLengthA < sideLengthC {
				base = sideLengthA
				if sideLengthB < sideLengthC {
					height = sideLengthB
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				} else {
					height = sideLengthC
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				}
			} else if sideLengthB < sideLengthA && sideLengthB < sideLengthC {
				base = sideLengthB
				if sideLengthA < sideLengthC {
					height = sideLengthA
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				} else {
					height = sideLengthC
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				}
			} else {
				base = sideLengthC
				if sideLengthA < sideLengthB {
					height = sideLengthA
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				} else {
					height = sideLengthB
					area := (base * height) / 2

					fmt.Printf("the area of right triangle is: %.2f", area)
				}
			}

		case acuteTriangle:
			if sideLengthA < sideLengthB && sideLengthA < sideLengthC {
				base = sideLengthA
				height = math.Sqrt((sideLengthC * sideLengthC) - (base / 2 * base / 2))
				area := (base * height) / 2

				fmt.Printf("the area of acute triangle is: %.2f", area)
			} else if sideLengthB < sideLengthA && sideLengthB < sideLengthC {
				base = sideLengthB
				height = math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))
				area := (base * height) / 2

				fmt.Printf("the area of acute triangle is: %.2f", area)
			} else {
				base = sideLengthC
				height = math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))
				area := (base * height) / 2

				fmt.Printf("the area of acute triangle is: %.2f", area)
			}

		case obtuseTriangle:

			maxAngle := math.Max(calculateAngle(sideLengthA, sideLengthB, sideLengthC), math.Max(calculateAngle(sideLengthA, sideLengthB, sideLengthC), calculateAngle(sideLengthA, sideLengthB, sideLengthC)))

			var baseSide, heightSide float64

			if maxAngle == calculateAngle(sideLengthA, sideLengthB, sideLengthC) {
				baseSide = sideLengthA
				heightSide = sideLengthB
				height = heightSide * math.Sin(maxAngle*math.Pi/180)

				area := (baseSide * height) / 2

				fmt.Printf("the area of obtuse triangle is: %.2f", area)
			} else if maxAngle == calculateAngle(sideLengthA, sideLengthB, sideLengthC) {
				baseSide = sideLengthB
				heightSide = sideLengthA
				height = heightSide * math.Sin(maxAngle*math.Pi/180)

				area := (baseSide * height) / 2

				fmt.Printf("the area of obtuse triangle is: %.2f", area)
			} else {
				baseSide = sideLengthC
				heightSide = sideLengthA
				height = heightSide * math.Sin(maxAngle*math.Pi/180)

				area := (baseSide * height) / 2

				fmt.Printf("the area of obtuse triangle is: %.2f", area)
			}

		default:
			fmt.Println("this option not is valid")
		}
	}
}

func runExercise10() {
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

	fmt.Println("exercise 9: Calculate the area of a triangle.")
	runExercise9()

	fmt.Println("exercise 10: Swap the values of two variables without using a third variable.")
	runExercise10()
}
