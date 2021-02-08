package main

import (
	"fmt"

	"github.com/ipr0/golang/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Trisha",
		LastName:  "McMillan",
	}
	fmt.Println(u)
	// go run .
	//go run github.com/ipr0/golang - fully qualified invocation
}
