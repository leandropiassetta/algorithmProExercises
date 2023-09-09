package exercise23

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunExercise23() {
	var input string
	var numbers []int

	fmt.Println("Enter a sequence of numbers ")

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

	average := sum / len(numbers)

	fmt.Printf("The average of the numbers entered is: %d", average)
}
