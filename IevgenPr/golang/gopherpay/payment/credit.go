// Package payment just a  docstring
package payment

import "fmt"

// CreditAccount blah
type CreditAccount1 struct{}

func (c *CreditAccount1) processPayment(amount float32) {
	fmt.Println("Processing credit transaction")
}

//CrCreditAccount uses named channel to accept amounts to charge
func CrCreditAccount(chargeCh chan float32) *CreditAccount1 {
	creditAccount := &CreditAccount1{}
	//// go routine listening to changes on that channel
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)
	////
	return creditAccount
}

//func main{}
