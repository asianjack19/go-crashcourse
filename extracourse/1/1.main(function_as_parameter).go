package main

import (
	"fmt"
)

//variabel filter line 8, memiliki tipe function dengan parameter string dan return value string
func sayHelloWithFilter(name string, filter func(string) string) {

	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
}
