package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	var address1 Address = Address{"Subang", "Jawa Barat", ""}
	var address2 *Address = &address1
	var address3 *Address = &address1
	//Jangan ganti

	address2.City = "Jakarta"
	*address2 = Address{"Tangerang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Papua",
	}
	//var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
