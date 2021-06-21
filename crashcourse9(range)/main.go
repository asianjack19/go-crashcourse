package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	//Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "jenna": "jenna@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k, v := range emails {
		fmt.Println("Name: " + k + "| Email: " + v)
	}
}
