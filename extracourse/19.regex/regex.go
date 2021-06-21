package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("M([a-z])([a-z])([a-z])([a-z])([a-z])([a-z])")

	fmt.Println(regex.MatchString("Michael"))
	fmt.Println(regex.MatchString("MiChael"))

	search := regex.FindAllString("Michael Michaal miChael", -1)

	fmt.Println(search)
}
