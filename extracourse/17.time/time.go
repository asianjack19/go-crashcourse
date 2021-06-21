package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())

	customDate := time.Date(2021, 10, 31, 24, 45, 60, 20, time.UTC)
	log.Println(customDate)

	// var layout string

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1945-06-05")
	fmt.Println(parse)
}
