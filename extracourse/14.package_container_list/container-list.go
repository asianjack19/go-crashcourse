package main

import (
	"container/list"
	"fmt"
)

//List itu double linked list di golang
func main() {
	data := list.New()

	data.PushBack("Michael")
	data.PushBack("Jack")
	data.PushBack("Chris")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
