package main

import (
	"fmt"
	"strings"
)

var (
	isDisabled, isSet bool
)

func main() {
	name := "beautiful go"
	author := "smart guy"

	fmt.Println(convert(name, author))

	bestScore := bestLeagueScore(13, 14, 2, 5, 8, 1)
	fmt.Println("Best score:", bestScore)
	fmt.Println(isSet, isDisabled)
	setVarsDefaults()
	fmt.Println(isSet, isDisabled)
	//https://tour.golang.org/basics/7
	x1, y1 := namedVars()
	fmt.Println(x1, y1)

}

func convert(s1, s2 string) (st1, st2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}

func bestLeagueScore(finishes ...int) int {
	best := finishes[0]
	for _, score := range finishes {
		if score > best {
			best = score
		}
	}
	return best
}

func setVarsDefaults() {
	isDisabled = true
	isSet = false
}

func namedVars() (x, y int) {
	x = 10
	y = 5
	return
}
