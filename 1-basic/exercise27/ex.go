package exercise27

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func quickSort(numbers []int) []int {
	// aqui é o caso base da recursão (quando a lista tem 1 ou 0 elementos), aqui vai parar a recursão porque não tem mais o que dividir a lista em duas partes menores para ordenar (a lista já está ordenada)
	if len(numbers) <= 1 {
		return numbers
	}

	var less, greater []int
	pivot := numbers[0]

	for _, number := range numbers[1:] {
		if number <= pivot {
			less = append(less, number)
		} else {
			greater = append(greater, number)
		}
	}

	// a
	less = quickSort(less)
	greater = quickSort(greater)

	return append(append(less, pivot), greater...)
}

func RunExercise27() {
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

	// for i := range numbers {
	// 	if numbers[i] != ascedingSortedOrder[i] {
	// 		fmt.Printf("Not sorted in ascending order: %v\n", numbers)
	// 	} else {
	// 		fmt.Printf("Sorted in ascending order: %v\n", numbers)
	// 	}
	// }
}
