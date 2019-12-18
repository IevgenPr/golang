package main

import (
	"github.com/gopherpay/payment"
)

func main() {
	var option payment.PaymentOption
	option = payment.CreateCreditAccount("Debra Ingram", "1111-2222-3333-4444", 8, 2024, 123)
	option.ProcessPayment(500)
	option = payment.CreateCashAccount()
	option.ProcessPayment(50)
	// fmt.Printf("Owner name: %s\n", option.OwnerName())
	// fmt.Printf("Card Number: %s\n", option.CardNumber())
	// fmt.Println("Trying to change card number")
	// err := option.SetCardNumber("invalid")
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }

	// fmt.Printf("Available Credit: %f\n", credit.AvailableCredit())

}
