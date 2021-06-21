package helper

import "fmt"

var Application = 1

func SayHello(name string) {
	fmt.Println("Hello", name)
}

//di go, access modifier(public, private) adanya pake huruf kapital
// contohnya pas declare var Application. Dia bisa diakses diluar package ini karena dia Kapital
