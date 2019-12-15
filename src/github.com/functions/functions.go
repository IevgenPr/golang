package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "beautiful go"
	author := "smart guy"

	fmt.Println(convert(name, author))
}

func convert(name, author string) (s1, s2 string) {
	name = strings.Title(name)
	author = strings.ToUpper(author)
	return name, author
}
