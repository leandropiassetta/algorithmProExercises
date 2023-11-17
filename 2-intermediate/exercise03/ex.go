package exercise03

import "fmt"

// Stack -> LIFO (Last In First Out)

type Stack struct {
	Elements []interface{}
	Top      int
}

const (
	MAX_SIZE = 10
)

// create a new stack
func NewStack() *Stack {
	return &Stack{
		Elements: make([]interface{}, 0, MAX_SIZE),
		Top:      -1,
	}
}

// adding operations to the stack:

// insert an element in the first position of the stack
func (stack *Stack) Push(element interface{}) {
	if stack.IsFull() {
		fmt.Println("Stack is full, don't possible insert new element")
		return
	}

	stack.Top++

	if stack.Top == len(stack.Elements) {
		stack.Elements = append(stack.Elements, element)
		return
	}

	stack.Elements[stack.Top] = element
}

// remove and return the first element of the stack if him don't empty
func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		fmt.Println("Stack is empty")
		return nil
	}

	element := stack.Elements[stack.Top]
	stack.Top--

	return element
}

func (stack *Stack) IsFull() bool {
	return len(stack.Elements) == MAX_SIZE

}

func (stack *Stack) IsEmpty() bool {
	return stack.Top < 0
}

func (stack *Stack) IsTop() interface{} {
	if stack.IsEmpty() {
		fmt.Println("Stack is empty")
		return nil
	}

	return stack.Elements[stack.Top]
}

func RunExercise03() {

	elementsInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	books := []string{"1984", "Romeu e julieta", "1984", "A cabana"}

	stackInt := NewStack()
	stackStr := NewStack()

	for _, element := range elementsInt {
		stackInt.Push(element)
	}

	for _, book := range books {
		stackStr.Push(book)
	}

	fmt.Printf("All elements of Ints: %v\n", stackInt.Elements)
	fmt.Printf("All elements of Books: %v\n", stackStr.Elements)

	// Stack -> LIFO (Last In First Out)

	fmt.Printf("First element of stack Ints : %d \n", stackInt.Elements[0])
	fmt.Printf("First element of stack Books : %s \n", stackStr.Elements[0])

	fmt.Printf("Last element of stack Ints : %d \n", stackInt.Elements[len(stackInt.Elements)-1])
	fmt.Printf("Last element of stack Books : %s \n", stackStr.Elements[len(stackStr.Elements)-1])

	fmt.Println("Stack of integers pop:")
	fmt.Println(stackInt.Pop())

	fmt.Println("Stack of strings pop:")
	fmt.Println(stackStr.Pop())

}
