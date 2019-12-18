package payment

import "fmt"

// go hybrid composition

//CreditAccount compositon of account
type CreditAccount struct{}

type CheckingAccount struct{}

type HybirdAccount struct {
	// no need to define a variable, all methods are transparently available
	CreditAccount //Account is now embedded into the structure
	CheckingAccount
}

// AvailableFunds specific to credit account
func (*CreditAccount) AvailableFunds() float32 {
	fmt.Println("Listing available funds: ")
	return 0
}

// AvailableFunds specific for checking account
func (*CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Listing available funds: ")
	return 0
}

//AvailableFunds specific for hybrid account - otherwise compilation fails
func (h *HybirdAccount) AvailableFunds() float32 {
	return h.CreditAccount.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}

func blah() {
	ha := &HybirdAccount{}
	fmt.Println(ha.AvailableFunds())
}
