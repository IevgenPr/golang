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

	chargeCh := make(chan float32)
	payment.CrCreditAccount(chargeCh)
	chargeCh <- 500
	// fmt.Println(account)

}
