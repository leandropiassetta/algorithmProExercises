package exercise17

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunExercise17() {
	var str string
	var reverseString string

	fmt.Println("Enter a string: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}

	strSplit := strings.Split(str, " ")
	strJoin := strings.Join(strSplit, "")

	for i := len(strJoin) - 1; i > -1; i-- {
		reverseString += string(strJoin[i])
	}

	strToLower := strings.ToLower(strJoin)
	reverseStringToLower := strings.ToLower(reverseString)

	if strToLower == reverseStringToLower {
		fmt.Printf("This string: %s is a palindrome!\n", str)
	} else {
		fmt.Printf("This string: %s not is a palindrome!\n", str)
	}
}
