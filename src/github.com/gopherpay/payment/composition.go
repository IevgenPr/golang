package payment

import "fmt"

// go composition

type Account struct{}

func (Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds: ")
	return 0
}

func (Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

//CreditAccount compositon of account
type CreditAccount struct {
	// no need to define a variable, all methods are transparently available
	Account
}

func blah() {
	ca := &CreditAccount{}
	funds := ca.AvailableFunds()
}
