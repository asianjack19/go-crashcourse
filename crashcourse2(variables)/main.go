package main

import "fmt"

func main() {

	//var
	//var name string = "Michael"
	// var name = "Michael"
	var age int = 31
	// const isCool=true
	var isCool = true
	isCool = false

	//shorthand
	// name := "Michael"
	size := 1.3
	name, email := "Michael", "michaeladriel08080@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", email)
}
