package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	// firstname string
	// lastname  string
	// city      string
	// gender    string
	// age       int

	firstname, lastname, city, gender string
	age                               int
}

//Greeting method (value receiver)
func (person Person) greet() string {
	return "Hello my name is" + person.firstname + person.lastname + " age: " + strconv.Itoa(person.age)
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age += 10
}

//getMarried method (pointer reciever)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastname = spouseLastName
	}
}

func main() {

	//Init person using struct
	person1 := Person{firstname: "hinata", lastname: "hyuga", city: "konoha", gender: "female", age: 32}

	//Alternative way
	person2 := Person{"James", "Jameson", "Tangerang", "male", 12}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstname)

	person1.hasBirthday()
	person2.getMarried("Jude")

	fmt.Println(person2.greet())

}
