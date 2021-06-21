package main

import "fmt"

func main() {
	x := 11
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n\n", y, x)
	}

	color := "blues"

	// if color == "red" {
	// 	fmt.Println("color is ", color)
	// } else if color == "blue" {
	// 	fmt.Println("color is ", color)
	// }

	switch color {
	case "red":
		{
			fmt.Println("color is ", color)
		}
	case "blue":
		{
			fmt.Println("color is ", color)
		}
	default:
		{
			fmt.Println("color unrecognized")
		}
	}

}
