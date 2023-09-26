package exercise09

import (
	"fmt"
	"math"
	"strings"
)

func calculateTriangleArea(a, b, c float64) float64 {
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
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

func RunExercise09() string {
	var sideLengthA float64
	var sideLengthB float64
	var sideLengthC float64
	var base float64
	var height float64

	fmt.Print("Enter the length of the first side of the triangle in cm: ")
	fmt.Scan(&sideLengthA)

	fmt.Print("Enter the length of the second side of the triangle in cm: ")
	fmt.Scan(&sideLengthB)

	fmt.Print("Enter the length of the third side of the triangle in cm: ")
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
	fmt.Println(options)

	var option string
	fmt.Scanf("%s\n", &option)

	switch option {
	case "a":
		fmt.Println("You have chosen to calculate the area of an equilateral triangle")
	case "b":
		fmt.Println("You have chosen to calculate the area of a scalene triangle")
	case "c":
		fmt.Println("You have chosen to calculate the area of an isosceles triangle")
	case "d":
		fmt.Println("You have chosen to calculate the area of a right triangle")
	case "e":
		fmt.Println("You have chosen to calculate the area of an acute triangle")
	case "f":
		fmt.Println("You have chosen to calculate the area of an obtuse triangle")
	default:
		return "this option not is valid please choose another option the options are: a - Equilateral triangle, b - Scalene triangle, c - Isosceles triangle, d - Right triangle, e - Acute triangle, f - Obtuse triangle"
	}

	fmt.Println("this sides chosen are: ", sideLengthA, sideLengthB, sideLengthC)

	isTriangleValid := sideLengthA+sideLengthB > sideLengthC && sideLengthA+sideLengthC > sideLengthB && sideLengthB+sideLengthC > sideLengthA

	if !isTriangleValid {
		return "the chosen sides do not form a triangle, please choose other sides"
	}

	fmt.Println("this sides form a triangle valid!!")

	// equilateral triangle has all sides equal
	equilateralTriangleIsValid := sideLengthA == sideLengthB && sideLengthB == sideLengthC
	equilateralTriangle := option == strings.ToLower("a") && equilateralTriangleIsValid

	if !equilateralTriangle && option == strings.ToLower("a") {
		return "this triangle is not equilateral because all sides are not equal, please choose side lengths that are equal or choose another option"
	}

	// scalene triangle has all sides different
	scaleneTriangleIsValid := sideLengthA != sideLengthB && sideLengthB != sideLengthC && sideLengthA != sideLengthC
	scaleneTriangle := option == strings.ToLower("b") && scaleneTriangleIsValid

	if !scaleneTriangle && option == strings.ToLower("b") {
		return "this triangle is not scalene because all sides are not different, please choose side lengths that are different or choose another option"
	}

	// isosceles triangle has two sides equal and one different side  (the different side is the base) and two equal angles (the angles opposite the equal sides) and the sum of the squares of the lengths of the legs is equal to the square of the length of the hypotenuse (Pythagorean theorem)

	isoscelesTriangleValid := sideLengthA == sideLengthB && sideLengthB != sideLengthC && sideLengthA != sideLengthC || sideLengthA == sideLengthC && sideLengthC != sideLengthB && sideLengthA != sideLengthB || sideLengthB == sideLengthC && sideLengthC != sideLengthA && sideLengthB != sideLengthA
	isoscelesTriangle := option == strings.ToLower("c") && isoscelesTriangleValid

	if !isoscelesTriangle && option == strings.ToLower("c") {
		return "this triangle is not isosceles because two sides are not equal and one different side, please choose side that fulfill this condition or choose another option"
	}

	// right triangle has one angle equal to 90 degrees and the sum of the squares of the lengths of the legs is equal to the square of the length of the hypotenuse (Pythagorean theorem)

	rightTriangleIsValid := math.Pow(sideLengthC, 2) == math.Pow(sideLengthA, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthA, 2) == math.Pow(sideLengthC, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthB, 2) == math.Pow(sideLengthA, 2)+math.Pow(sideLengthC, 2)
	rightTriangle := option == strings.ToLower("d") && rightTriangleIsValid

	if !rightTriangle && option == strings.ToLower("d") {
		return "this triangle is not right because one side is not equal that sum of the squares of the lengths of the legs, please choose side that fulfill this condition or choose another option"
	}
	// acute triangle has all angles less than 90 degrees and the sum of the squares of the lengths of the legs is greater than the square of the length of the hypotenuse (Pythagorean theorem)

	acuteTriangleIsValid := math.Pow(sideLengthC, 2) < math.Pow(sideLengthA, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthA, 2) < math.Pow(sideLengthC, 2)+math.Pow(sideLengthB, 2) || math.Pow(sideLengthB, 2) < math.Pow(sideLengthA, 2)+math.Pow(sideLengthC, 2)
	acuteTriangle := (option == "e") && acuteTriangleIsValid

	if !acuteTriangle && option == strings.ToLower("e") {
		return "this triangle is not acute because one side is not less than the sum of the squares of the lengths of the legs, please choose side that fulfill this condition or choose another option"
	}

	// obtuse triangle has one angle greater than 90 degrees and the sum of the squares of the lengths of the legs is less than the square of the length of the hypotenuse (Pythagorean theorem)

	obtuseTriangle := option == strings.ToLower("f") && checkObtuse(sideLengthA, sideLengthB, sideLengthC)

	if !checkObtuse(sideLengthA, sideLengthB, sideLengthC) && option == strings.ToLower("f") {
		return "this triangle is not obtuse because one angle is not greater than 90 degrees, please choose side that fulfill this condition or choose another option"
	}

	switch {

	case equilateralTriangle:

		height = (math.Sqrt(3) / 2) * sideLengthA

		base = sideLengthC

		area := (base * height) / 2

		return fmt.Sprintf("the area of equilateral triangle is: %.3fcm²", area)

	case scaleneTriangle:
		perimeter := sideLengthA + sideLengthB + sideLengthC
		semiPerimeter := perimeter / 2

		area := math.Sqrt(semiPerimeter * (semiPerimeter - sideLengthA) * (semiPerimeter - sideLengthB) * (semiPerimeter - sideLengthC))

		return fmt.Sprintf("the area of scalene triangle is: %.3fcm²", area)
	case isoscelesTriangle:

		if sideLengthC == sideLengthB {
			base = sideLengthA
			// theorem of pythagoras
			height := math.Sqrt((sideLengthC * sideLengthC) - (base / 2 * base / 2))

			area := (base * height) / 2

			return fmt.Sprintf("the area of isosceles triangle is: %.3fcm²", area)
		} else if sideLengthA == sideLengthC {
			base = sideLengthB
			// theorem of pythagoras
			height := math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))

			area := (base * height) / 2

			return fmt.Sprintf("the area of isosceles triangle is: %.3fcm²", area)
		} else {
			base = sideLengthC
			// theorem of pythagoras
			height := math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))

			area := (base * height) / 2

			return fmt.Sprintf("the area of isosceles triangle is: %.3fcm²", area)
		}

	case rightTriangle:
		if sideLengthA < sideLengthB && sideLengthA < sideLengthC {
			base = sideLengthA
			if sideLengthB < sideLengthC {
				height = sideLengthB
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			} else {
				height = sideLengthC
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			}
		} else if sideLengthB < sideLengthA && sideLengthB < sideLengthC {
			base = sideLengthB
			if sideLengthA < sideLengthC {
				height = sideLengthA
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			} else {
				height = sideLengthC
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			}
		} else {
			base = sideLengthC
			if sideLengthA < sideLengthB {
				height = sideLengthA
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			} else {
				height = sideLengthB
				area := (base * height) / 2

				return fmt.Sprintf("the area of right triangle is: %.3fcm²", area)
			}
		}

	case acuteTriangle:
		if sideLengthA < sideLengthB && sideLengthA < sideLengthC {
			base = sideLengthA
			height = math.Sqrt((sideLengthC * sideLengthC) - (base / 2 * base / 2))
			area := (base * height) / 2

			return fmt.Sprintf("the area of acute triangle is: %.3fcm²", area)
		} else if sideLengthB < sideLengthA && sideLengthB < sideLengthC {
			base = sideLengthB
			height = math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))
			area := (base * height) / 2

			return fmt.Sprintf("the area of acute triangle is: %.3fcm²", area)
		} else {
			base = sideLengthC
			height = math.Sqrt((sideLengthA * sideLengthA) - (base / 2 * base / 2))
			area := (base * height) / 2

			return fmt.Sprintf("the area of acute triangle is: %.3fcm²", area)
		}

	case obtuseTriangle:

		maxAngle := math.Max(calculateAngle(sideLengthA, sideLengthB, sideLengthC), math.Max(calculateAngle(sideLengthA, sideLengthB, sideLengthC), calculateAngle(sideLengthA, sideLengthB, sideLengthC)))

		var baseSide, heightSide float64

		if maxAngle == calculateAngle(sideLengthA, sideLengthB, sideLengthC) {
			baseSide = sideLengthA
			heightSide = sideLengthB
			height = heightSide * math.Sin(maxAngle*math.Pi/180)

			area := (baseSide * height) / 2

			return fmt.Sprintf("the area of obtuse triangle is: %.3fcm²", area)
		} else if maxAngle == calculateAngle(sideLengthA, sideLengthB, sideLengthC) {
			baseSide = sideLengthB
			heightSide = sideLengthA
			height = heightSide * math.Sin(maxAngle*math.Pi/180)

			area := (baseSide * height) / 2

			return fmt.Sprintf("the area of obtuse triangle is: %.3fcm²", area)
		} else {
			baseSide = sideLengthC
			heightSide = sideLengthA
			height = heightSide * math.Sin(maxAngle*math.Pi/180)

			area := (baseSide * height) / 2

			return fmt.Sprintf("the area of obtuse triangle is: %.3fcm²", area)
		}

	default:
		return "this option not is valid"
	}
}
