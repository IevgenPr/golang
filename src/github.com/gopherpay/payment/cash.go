// Package payment just a  docstring
package payment

import "fmt"

// Cash is just an empty struct
type Cash struct{}

// CreateCashAccount is sort of constructor which returns Cash
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment implements respective interface
func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing cash payment...")
	return true
}
