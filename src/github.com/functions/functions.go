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

func convert(s1, s2 string) (st1, st2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}
