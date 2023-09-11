package exercise26

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise26() {
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

	var sum int

	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("The sum of all numbers entered is: %d\n", sum)
}
