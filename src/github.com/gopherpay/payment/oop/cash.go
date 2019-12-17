package oop

// Cash blah
type Cash struct{}

// CreateCashAccount blah
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment blah
func (c Cash) ProcessPayment(amount float32) bool {
	return true
}
