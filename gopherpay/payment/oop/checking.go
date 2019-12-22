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
		balance:       500.0, // should come from a web service
	}
}

// ProcessPayment blah
func (c CheckingAccount) ProcessPayment(amount float32) bool {
	return true
}

// Balance blah
func (c CheckingAccount) Balance() float32 {
	return c.balance
}
