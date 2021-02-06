package main

import (
	"fmt"
	"strings"

	"github.com/tdrgr/golang/gopherpay/payment/procedural"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	procedural.PayWithCash(amount)

	fmt.Println(strings.Repeat("-", 10))

	credit := &procedural.CreditCard{
		OwnerName:       "Debra Ingram",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 8,
		ExpirationYear:  2024,
		SecurityCode:    123,
		AvailableCredit: 5000,
	}
	fmt.Println("Paying with credit card")
	fmt.Printf("Initial balance: %.2f\n", credit.AvailableCredit)
	procedural.PayWithCredit(credit, amount)
	fmt.Printf("Balance now:  %.2f\n", credit.AvailableCredit)

	fmt.Println(strings.Repeat("-", 10))

	checkingAccount := &procedural.CheckingAccount{
		AccountOwner:  "Debra Ingram",
		RoutingNumber: "2359328048",
		AccountNumber: "234305555",
		Balance:       600,
	}

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: %.2f\n", checkingAccount.Balance)
	procedural.PayWithCheck(checkingAccount, amount)
	fmt.Printf("Balance now:  %.2f\n", checkingAccount.Balance)

	fmt.Println(strings.Repeat("-", 10))

}
