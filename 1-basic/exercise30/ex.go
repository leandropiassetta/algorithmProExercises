package exercise30

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise30() {
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

	var arrayWithoutDuplicates []int

	unique := make(map[int]bool)

	for _, number := range numbers {
		if !unique[number] {
			unique[number] = true
			arrayWithoutDuplicates = append(arrayWithoutDuplicates, number)
		}
	}

	fmt.Println("The sequence without duplicates is: ", arrayWithoutDuplicates)
}
