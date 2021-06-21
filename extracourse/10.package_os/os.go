package main

import (
	_ "database/sql"
	"fmt"
	"os"
)

func main() {
	// args := os.Args
	// fmt.Println(args)

	name, error := os.Hostname()

	if error == nil {
		fmt.Println("Hostname: ", name)
	} else {
		fmt.Println(error)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if password == "" {
		fmt.Println("x")
	}

	fmt.Println(username, password)
}
