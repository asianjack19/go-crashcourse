package main

import "fmt"

func main() {
	//Arrays
	// var fruitArr [2]string

	//Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Watermelon"

	//Array, declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}
	// fmt.Println(fruitArr)

	//Slice, declare and assign
	fruitSlice := []string{"Apple", "Orange", "Watermelon"}

	fmt.Println(fruitSlice[0:2])
}
