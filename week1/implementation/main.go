package main

import "fmt"

type Node struct {
	Next *Node
	Data any
}

type List struct {
	Head *Node
}

func NewNode(datas any) *Node {
	return &Node{Data: datas}
}

type LinkedList interface {
	Add(datas any)
	Get(any)
	GetAll()
	Remove(any)
}

func NewList() LinkedList {
	L := List{
		Head: nil,
	}

	return &L
}

func (L *List) Add(datas any) {
	if L.Head == nil {
		L.Head = NewNode(datas)
	} else {
		current := L.Head
		for {
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
		current.Next = NewNode(datas)
	}
}

func (L *List) GetAll() {
	if L.Head == nil {
		fmt.Println("list empty")
		return
	}

	current := L.Head
	for {
		fmt.Print(current.Data, " ")
		if current.Next != nil {
			current = current.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (L *List) Get(datas any) {
	if L.Head == nil {
		fmt.Println("list empty")
		return
	}

	current := L.Head
	for {
		if current.Next == nil && current.Data != datas {
			fmt.Println("data Not found")
			break
		}
		if current.Next != nil {
			if current.Data == datas {
				break
			}
			current = current.Next
		} else {
			break
		}
	}
}

func (L *List) Remove(datas any) {
	if L.Head == nil {
		fmt.Println("list empty")
		return
	}

	if L.Head.Data == datas {
		L.Head = L.Head.Next
		return
	}

	current := L.Head
	for current.Next != nil && current.Next.Data != datas {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func main() {
	ll := NewList()

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	ll.GetAll()

	ll.Remove(20)
	ll.Add(40)
	ll.GetAll()
}