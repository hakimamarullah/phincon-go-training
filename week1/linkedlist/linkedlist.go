package linkedlist

import "fmt"

type Node struct {
  data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

type Collections interface {
	Add(data interface{})
	Display()
	Length() int64
}

func (ll *LinkedList) Add(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	}else{
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func (ll *LinkedList) Length() int64 {
	current := ll.head
	counter := 0

	for current != nil {
		counter++
		current = current.next
	}
	return int64(counter)
}

func main() {
	
}