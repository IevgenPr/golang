package main

import "fmt"

func main() {
	if firstRank, secondRank := "678", "39"; firstRank < secondRank {
		fmt.Println("First is doing better")
	} else if firstRank == secondRank {
		fmt.Println("First is equals" +
			" to second")
	} else {
		fmt.Println("Second is doing better")
	}
}
