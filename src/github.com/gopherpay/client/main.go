package main

import (
	"fmt"
	"github.com/gopherpay/payment"
)

func main()  {
	credit := payment.CreateCreditAccount("Debra Ingram","1111-2222-3333-4444",8,2024,123)
	
	fmt.Printf("Owner name: %s\n", credit.OwnerName())
	fmt.Printf("Card Number: %s\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("invalid")
	if err != nil{
		fmt.Printf("Error: %v\n", err)
	}

fmt.Printf("Available Credit: %f\n", credit.AvailableCredit())

}