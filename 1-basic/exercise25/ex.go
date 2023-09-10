package exercise25

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise25() {
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

	minNumber := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if minNumber < numbers[i] {
			continue
		} else {
			minNumber = numbers[i]
		}
	}

	fmt.Printf("The minimium number is: %v\n", minNumber)
}
