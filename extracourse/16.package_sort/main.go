package main

import (
	"log"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

//Struct method
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{
			Name: "Michael",
			Age:  40,
		},
		{
			Name: "Jack",
			Age:  30,
		},
		{
			Name: "John",
			Age:  70,
		},
		{
			Name: "James",
			Age:  90,
		},
	}
	//Pakai package sort, lalu panggil function sort
	sort.Sort(UserSlice(users))

	log.Println(users)

}
