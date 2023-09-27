package exercise01

import "fmt"

type Person struct {
	Name     string
	Lastname string
	Age      int
}

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value Person
	Next  *Node
}

func (list *List) Append(person Person) {
	node := &Node{Value: person}

	if list.Head == nil {
		list.Head = node
	}

	if list.Tail != nil {
		list.Tail.Next = node
	}

	list.Tail = node
}

func (list *List) Search(name string) (person Person) {
	node := list.Head

	for node != nil {
		if node.Value.Name == name {
			person = node.Value
			break
		}
		node = node.Next
	}

	if node != nil {
		return node.Value
	}

	return
}

func (list *List) Delete(name string) {
	if list.Head.Value.Name == name {
		list.Head = list.Head.Next
		return
	}

	previous := list.Head
	node := list.Head.Next

	for node != nil {
		if node.Value.Name == name {
			previous.Next = node.Next
			break
		}

		previous = node
		node = node.Next
	}

	if list.Tail == node {
		list.Tail = previous
	}
}

func (list *List) Display() {
	node := list.Head

	for node != nil {
		fmt.Println("name" + node.Value.Name)
		node = node.Next
	}
}

func RunExercise01() {
	list := List{}

	mickey := Person{"Mickey", "Mouse", 100}
	pato := Person{"Pato", "Donald", 93}
	pateta := Person{"Goofy", "Goof", 99}
	tioPatinhas := Person{"Tio", "Patinhas", 100}

	list.Append(mickey)
	list.Append(pato)
	list.Append(pateta)
	list.Append(tioPatinhas)

	list.Display()

	fmt.Println("------------------------")

	person := list.Search("Pato")

	fmt.Println(person)

	fmt.Println("------------------------")
	list.Delete(person.Name)

	list.Display()
}
