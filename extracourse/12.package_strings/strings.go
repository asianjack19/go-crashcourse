package main

import (
	"fmt"
	"strings"
)

func main() {

	//Contains return boolean
	fmt.Println(strings.Contains("Michael Adriel", "Michael"))

	//Split return splice
	slice := []string{}
	slice = strings.Split("Michael Adriel", " ")
	fmt.Println(slice[0])

	//strings.ToLower => convert lowercase
	//strings.ToUpper => convert uppercase

	//Trim remove kanan kiri
	fmt.Println(strings.Trim("..........................MIchael Adriel Darmawan........................", "."))

	//Replace all ubah semua
	fmt.Println(strings.ReplaceAll("Michael Michael Adriel Darmawan", "Michael", "Ekel"))
}
