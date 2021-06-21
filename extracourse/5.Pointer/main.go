package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {
	//Cara 1 bikin pointer
	// address1 := Address{"Subang", "Jawa Barat", ""}
	// address2 := &address1

	//Cara 2 bikin pointer
	var address1 Address = Address{"Subang", "Jawa Barat", ""}
	var address2 *Address = &address1
	var address3 *Address = &address1
	//Jangan ganti

	address2.City = "Jakarta"

	//&Address bikin memori baru buat address 2, address 1 masih tetep refer ke alamat awal
	// address2 = &Address{"Tangerang", "Jawa Timur", "Indonesia"}

	//Pake bintang jadi semua yang refer ke address2 valuenya keganti semu
	*address2 = Address{"Tangerang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	//Bikin pointer kosong
	var address4 *Address = new(Address)
	fmt.Println(address4)
}
