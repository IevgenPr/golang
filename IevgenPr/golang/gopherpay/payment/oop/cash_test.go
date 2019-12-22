package oop

//go test github.com/gopherpay/payment/oop -v
import (
	"reflect"
	"testing"
)

func TestCreateCashAccount(t *testing.T) {
	want := reflect.TypeOf(new(Cash))
	got := reflect.TypeOf(CreateCashAccount())
	if got != want {
		t.Errorf("PayWithCash = %s; want %s", got, want)
	}
}
func TestProcessPayment(t *testing.T) {
	// set up
	account := CreateCashAccount()
	got := account.ProcessPayment(100)
	if !got {
		t.Errorf("ProcessPayment = %t; want true", got)
	}

	// t.Run("Credit 100 from 5000, get 4900 in available credit",
	// 	func(t *testing.T) {
	// 		const remainingCredit = 4400
	// 		got := PayWithCredit(credit, amount)
	// 		if !got {
	// 			t.Errorf("PayWithCheck = %t; want true", got)
	// 		}
	// 		if credit.AvailableCredit != remainingCredit {
	// 			t.Errorf("Available credit = %.2f; want %d",
	// 				credit.AvailableCredit, remainingCredit)
	// 		}
	// 	})

	// t.Run("Check failed to withdraw",
	// 	func(t *testing.T) {
	// 		got := PayWithCredit(credit, 5000)
	// 		if got {
	// 			t.Errorf("PayWithCredit = %t; want false", got)
	// 		}
	// 		if credit.AvailableCredit != 4400 {
	// 			t.Errorf("PayWithCredit = %f; want 4400", credit.AvailableCredit)
	// 		}
	// 	})
}
