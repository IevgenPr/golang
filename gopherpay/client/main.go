package main

import (
	"github.com/tdrgr/golang/gopherpay/payment"
)

// PaymentOption is a common interface for various options.
// CreditCard and Cash structs implement method of this interface
type PaymentOption interface {
	ProcessPayment(float32) bool
}

func main() {

	var option PaymentOption
	option = payment.CreateCreditAccount("Debra Ingram", "1111-2222-3333-4444", 8, 2024, 123)
	option.ProcessPayment(500)
	option = payment.CreateCashAccount()
	option.ProcessPayment(50)

	chargeCh := make(chan float32)
	payment.CrCreditAccount(chargeCh)
	chargeCh <- 500
	// fmt.Println(account)

}
