package main

import "fmt"

func main() {
	//define map
	// emails := make(map[string]stringmap[string]string)
	// emails := make(map[string]string)

	//assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["mike"] = "mike@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "jenna": "jenna@gmail.com"}

	fmt.Println(emails)

	//delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
