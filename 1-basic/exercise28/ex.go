package exercise28

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise28() {
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

	fmt.Println("Enter a number to count the occurrences in the sequence: ")
	fmt.Scan(&number)

	var count int

	for _, element := range numbers {
		if element == number {
			count++
		}
	}

	fmt.Printf("The count of occurrences of %d in the sequence is: %d\n", number, count)
}
