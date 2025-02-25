package main

import (
	"container/list"
	"fmt"
)

// Just learning how to use go's list
func main() {
	// Creating a new list and put stuff on it
	my_list := list.New()

	element4 := my_list.PushBack(4)
	element1 := my_list.PushFront(1)

	my_list.InsertBefore(3, element4)
	my_list.InsertAfter(2, element1)

	// print list by iterating
	for element := my_list.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)

	}
}
