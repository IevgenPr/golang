package main

import "fmt"

func main() {
	firstRank := "678"
	secondRank := "39"
	if firstRank < secondRank {
		fmt.Println("First is doing better")
	} else if firstRank == secondRank {
		fmt.Println("First is equals" +
			" to second")
	} else {
		fmt.Println("Second is doing better")
	}
}
