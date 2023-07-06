package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type List interface {
	Add(data interface{})
	Get(data interface{}) (bool, *interface{})
	Display()
	Remove(data interface{})
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDoublyLinkedList() List {
	return &DoublyLinkedList{Head: nil, Tail: nil}
}

func NewNode(data interface{}) *Node {
	return &Node{Data: data, Next: nil, Prev: nil}
}

func (L *DoublyLinkedList) Add(data interface{}) {
	newNode := NewNode(data)
	if L.Head == nil {
		L.Head = newNode
		L.Tail = newNode
		return
	}

	L.Tail.Next = newNode
	newNode.Prev = L.Tail
	L.Tail = newNode

}
func (L *DoublyLinkedList) Get(data interface{}) (bool, *interface{}) {
	if L.Head == nil {
		return false, nil
	}

	current := L.Head
	for current != nil && current.Data != data {
		current = current.Next
	}

	if current != nil {
		return true, &current.Data
	}
	return false, nil
}


func (L *DoublyLinkedList) Display() {
	current := L.Head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}


func (L *DoublyLinkedList) Remove(data interface{}) {
	if L.Head == nil {
		return
	}

	if L.Head.Data == data {
		L.Head = L.Head.Next
		L.Head.Prev = nil
		return
	}

	current := L.Head.Next
	for current != nil && current.Data != data {
		current = current.Next
	}
	
	if current != nil {
		if current.Next == nil {
			current.Prev.Next = nil
			L.Tail = current.Prev
		}else{
			current.Next.Prev = current.Prev
			current.Prev.Next = current.Next
		}
	}
}

func main() {
	var list List = NewDoublyLinkedList()

	var cmd int
	var data int
	for {
		fmt.Printf("Enter Operation\n1: Add\n2: Get\n3: Remove\n0: Exit\nCommand: ")
		fmt.Scan(&cmd)

		if cmd == 1 {
			fmt.Print("Enter data: ")
			fmt.Scan(&data)
			list.Add(data)
			fmt.Print("Current list: ")
			list.Display()
		}else if cmd == 2{
			fmt.Print("Enter data to find: ")
			fmt.Scan(&data)
			fmt.Println()
			find, _ := list.Get(data)
			fmt.Printf("FOUND: %v\n", find)
		}else if cmd == 3 {
			fmt.Print("Enter data to remove: ")
			fmt.Scan(&data)
			fmt.Println()
			list.Remove(data)
			fmt.Print("Current list: ")
			list.Display()
		}else if cmd == 0 {
			fmt.Println("Exited! DoublyLinkedList!")
			break
		}else {
			fmt.Println("Commmand Not Found")
			break
		}
		fmt.Println("========================================================")
	}
}
