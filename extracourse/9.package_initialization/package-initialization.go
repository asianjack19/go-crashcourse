package main

import (
	"fmt"

	"github.com/asianjack19/go_crash_course/extracourse/9.package_initialization/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
