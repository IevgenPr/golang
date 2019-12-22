package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// CreditCard blah
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditAccount blah
func CreateCreditAccount(
	ownerName, cardNumber string,
	expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
		availableCredit: 5000.0,
	}
}

// ProcessPayment blah
func (CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

// OwnerName blah
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName ownerName setter
func (c CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = value
	return nil
}

// CardNumber blah
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber ownerName setter
func (c CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid card number format")
	}
	c.cardNumber = value
	return nil
}

// ExpirationDate blahs
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate blah
func (c CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() ||
		(year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Invalid expiration data")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode blahs
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode blahs
func (c CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("Invalid security code ")
	}
	c.securityCode = value
	return nil
}

// AvailableCredit blah
// blah1
// blah2
func (c CreditCard) AvailableCredit() float32 {
	return 5000.
}
