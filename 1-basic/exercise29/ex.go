package exercise29

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise29() {
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

	var number int

	fmt.Println("Enter a number to find its index in the sequence:")
	fmt.Scan(&number)

	var findIndexes []int

	for index, element := range numbers {
		if element == number {
			findIndexes = append(findIndexes, index)
		}
	}

	if len(findIndexes) == 0 {
		fmt.Printf("Not found indexes for this number %v, in the sequence %v\n", number, numbers)
	} else {
		fmt.Printf("The indexes for thins number %v found in the sequence %v are: %v\n", number, numbers, findIndexes)
	}
}
