package exercise49

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunExercise49() {
	fmt.Print("Enter a phrase: ")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the input.")
			return
		}

		if len(input) == 0 {
			fmt.Println("The input cannot be empty.")
			return
		}

		words := strings.Fields(input)

		firstLetter := string(words[0][0])
		isTautogram := true

		for i := 1; i < len(words); i++ {
			compareLetter := string(words[i][0])

			if firstLetter != compareLetter {
				isTautogram = false
				fmt.Printf("The input '%s' is not a tautogram\n", input)
				break
			}

			if isTautogram {
				fmt.Printf("The input '%s' is a tautogram\n", input)
			}
		}
	}
}
