package oop

// CheckingAccount blah
type CheckingAccount struct {
	accountOwner  string
	routingNumber string
	accountNumber string
	balance       float32
}

// CreateCheckingAccount blah
func CreateCheckingAccount(
	accountOwner, routingNumber, accountNumber string) *CheckingAccount {
	return &CheckingAccount{
		accountOwner:  accountOwner,
		routingNumber: routingNumber,
		accountNumber: accountNumber,
		balance:       0.0,
	}
}
