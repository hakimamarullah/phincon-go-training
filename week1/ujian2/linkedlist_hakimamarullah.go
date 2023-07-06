package main

import (
	"fmt"
	"strings"
	"time"
)


type Catalog struct {
	Id int64
	Name string
	Brand string
	Price int64
	Quantity int64
}

type ListCatalog interface {
	AddCatalog(data any)
	DisplayCatalog()
	DeleteByName(name string)
	FindById(id int64) (bool, *Catalog)
	FindByName(name string) (bool, *Catalog)
	UpdateName(id int64, name string)
}

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
}

func NewNode(data any) *Node {
	return &Node{Data: data, Next: nil}
}

func NewList() ListCatalog {
	return &List{Head: nil}
}

func (L *List) AddCatalog(data any) {
	if L.Head == nil {
		L.Head = NewNode(data)
	} else {
		current := L.Head
		for {
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
		current.Next = NewNode(data)
	}
}

func (L *List) DisplayCatalog() {
	if L.Head == nil {
		fmt.Println("No Catalog")
		return
	}

	current := L.Head
	for {
		fmt.Print(current.Data, ", ")
		if current.Next != nil {
			current = current.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (L *List) DeleteByName(name string) {
	if L.Head == nil {
		fmt.Println("No Catalog")
		return
	}

	if strings.EqualFold(L.Head.Data.(*Catalog).Name, name) {
		L.Head = L.Head.Next
		return
	}

	current := L.Head
	for current.Next != nil && !strings.EqualFold(current.Data.(*Catalog).Name, name) {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (L *List) FindById(id int64) (bool, *Catalog) {
	if L.Head == nil {
		return false, nil
	}

	if L.Head.Data.(*Catalog).Id == id {
		return true, L.Head.Data.(*Catalog)
	}

	current  := L.Head.Next
	for current != nil && current.Data.(*Catalog).Id == id {
		current = current.Next
	}

	if current != nil {
		return true, current.Data.(*Catalog)
	}

	return false, nil

}

func (L *List) FindByName(name string) (bool, *Catalog) {
	if L.Head == nil {
		return false, nil
	}

	if strings.EqualFold(L.Head.Data.(*Catalog).Name, name) {
		return true, L.Head.Data.(*Catalog)
	}

	current  := L.Head.Next
	for current != nil && strings.EqualFold(current.Data.(*Catalog).Name, name){
		current = current.Next
	}

	if current != nil {
		return true, current.Data.(*Catalog)
	}

	return false, nil
}

func (L *List) UpdateName(id int64, name string) {
	if find, data := L.FindById(id); find {
		data.Name = name
	}
}

func main() {
	id := time.Now().UnixNano()
	time.Sleep(time.Millisecond)
	id2 := time.Now().UnixNano()
	catalog1 := &Catalog{Id: id, Name: "Baju", Brand: "Uniqlo", Price: 200, Quantity: 100}
	catalog2 := &Catalog{Id: id2, Name: "Celana", Brand: "Uniqlo", Price: 200, Quantity: 100}
	var list ListCatalog = NewList()

	list.AddCatalog(catalog1)
	list.AddCatalog(catalog2)

	list.DisplayCatalog()
	list.DeleteByName("C")
	fmt.Println("After Delete")
	list.DisplayCatalog()

	fmt.Println("====================")
	list.DisplayCatalog()
	list.DeleteByName("Baju")
	fmt.Println("After Delete")
	list.DisplayCatalog()

	list.UpdateName(id2, "Jeans")
	if find, cat := list.FindByName("Jeans"); find {
		fmt.Println(cat)
	}

	list.DisplayCatalog()

	
}
