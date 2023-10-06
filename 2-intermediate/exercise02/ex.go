package exercise02

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []interface{}
}

// create a new queue
func NewQueue() *Queue {
	return &Queue{
		Elements: []interface{}{},
	}
}

// adding operations to the queue:

// insert an element in the last position of the queue
func (queue *Queue) Insert(element interface{}) {
	queue.Elements = append(queue.Elements, element)
}

// remove and return the first element of the queue if him don't empty
func (queue *Queue) Remove() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	element := queue.Elements[0]
	queue.Elements = queue.Elements[1:]

	return element, nil
}

// return an element of the queue without remove him, if the queue is empty return nil
func (queue *Queue) Peek() (interface{}, error) {
	if queue.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	return queue.Elements[0], nil
}

// return true if the queue is empty
func (queue *Queue) IsEmpty() bool {
	return len(queue.Elements) == 0
}

// return the size of the queue
func (queue *Queue) Length() int {
	return len(queue.Elements)
}

// remove all elements of the queue
func (queue *Queue) Clear() {
	queue.Elements = make([]interface{}, 0)
}

// print all elements of the queue
func (queue *Queue) Print() {
	for i, element := range queue.Elements {
		fmt.Printf("Position %d: %v\n", i, element)
	}

	fmt.Println()
}

// reverse the queue
func (queue *Queue) Reverse() {
	length := len(queue.Elements)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		queue.Elements[i], queue.Elements[j] = queue.Elements[j], queue.Elements[i]
	}
}

func (queue *Queue) Contains(element interface{}) bool {
	for _, queueElement := range queue.Elements {
		if queueElement == element {
			return true
		}
	}

	return false
}

func (queue *Queue) ToArray() []interface{} {
	copyElements := make([]interface{}, len(queue.Elements))
	copy(copyElements, queue.Elements)

	return copyElements
}

// return true if the queue is full
func (queue *Queue) IsFull() bool {
	return len(queue.Elements) == cap(queue.Elements)
}

func (queue *Queue) Cap() int {
	return cap(queue.Elements)
}

func RunExercise02() {
	elementsInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	books := []string{"1984", "Romeu e julieta", "1984", "A cabana"}

	queueInt := NewQueue()
	queueStr := NewQueue()

	for _, element := range elementsInt {
		queueInt.Insert(element)
	}

	for _, book := range books {
		queueStr.Insert(book)
	}

	fmt.Println("Queue of integers:")
	queueInt.Print()

	fmt.Println("Queue of strings:")
	queueStr.Print()

	fmt.Println("Queue of integers reversed:")
	queueInt.Reverse()
	queueInt.Print()

	fmt.Println("Queue of strings reversed:")
	queueStr.Reverse()
	queueStr.Print()

	fmt.Println("Queue of integers lenght:")
	fmt.Println(queueInt.Length())

	fmt.Println("Queue of strings lenght:")
	fmt.Println(queueStr.Length())

	fmt.Println("Queue of integers is empty?")
	fmt.Println(queueInt.IsEmpty())

	fmt.Println("Queue of strings is empty?")
	fmt.Println(queueStr.IsEmpty())

	fmt.Println("Queue of integers is full?")
	fmt.Println(queueInt.IsFull())

	fmt.Println("Queue of strings is full?")
	fmt.Println(queueStr.IsFull())

	fmt.Println("Queue of integers contains 5?")
	fmt.Println(queueInt.Contains(5))

	fmt.Println("Queue of strings contains '1984'?")
	fmt.Println(queueStr.Contains("1984"))

	fmt.Println("Queue of integers peek:")
	fmt.Println(queueInt.Peek())

	fmt.Println("Queue of strings peek:")
	fmt.Println(queueStr.Peek())

	fmt.Println("Queue of integers remove:")
	fmt.Println(queueInt.Remove())

	fmt.Println("Queue of strings remove:")
	fmt.Println(queueStr.Remove())

	fmt.Println("Queue of integers clear:")
	queueInt.Clear()

	fmt.Println("Queue of strings clear:")
	queueStr.Clear()
}
