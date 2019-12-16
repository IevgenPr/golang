package main

import (
	"fmt"
	"os"
)

func main() {
	//here open function returns error as last argument
	_, err := os.Open("somepath")
	// nil here is success
	if err != nil {
		fmt.Println("there was error opening file", err)
	}
}
