package main

import "fmt"

func main() {
	maps := make(map[[2]int]func(...interface{}))

	maps[[2]int{1,2}] = func (args ...interface{}) {
		fmt.Printf("Hello, %s", args[0])
	}

	maps[[2]int{1,2}]("Sally!")
}