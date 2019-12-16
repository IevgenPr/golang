package main

import (
	"fmt"
	"time"
)

func main() {
	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hi there")
	}() // parens make func self executing

	func() {
		fmt.Println("Niccce")
	}()
}
