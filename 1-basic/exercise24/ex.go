package exercise24

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise24() {
	var input string
	var numbers []int

	fmt.Println("Enter a sequence of numbers, to finish enter a non-number character: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		input = scanner.Text()

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Numbers entered: %v\n", numbers)
			break
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	maxNumber := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if maxNumber > numbers[i] {
			continue
		} else {
			maxNumber = numbers[i]
		}
	}

	fmt.Printf("The maximium number is: %v\n", maxNumber)
}
