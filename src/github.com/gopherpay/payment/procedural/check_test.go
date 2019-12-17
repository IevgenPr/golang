package procedural

//go test github.com/gopherpay/payment/procedural -v

import (
	"testing"
)

func TestPayWithCheck(t *testing.T) {
	//set up
	const amount = 500

	checkingAccount := &CheckingAccount{
		AccountOwner:  "Debra Ingram",
		RoutingNumber: "2359328048",
		AccountNumber: "234305555",
		Balance:       600,
	}

	//check withdrawal
	t.Run("withdraw 500 from my account with 600, new balance 100",
		func(t *testing.T) {
			got := PayWithCheck(checkingAccount, amount)
			if !got {
				t.Errorf("PayWithCheck = %t; want true", got)
			}
			if checkingAccount.Balance != 100 {
				t.Errorf("PayWithCheck.Balance = %f; want 100", checkingAccount.Balance)
			}
		})
	// some other test
	t.Run("Check failed to withdraw",
		func(t *testing.T) {
			got := PayWithCheck(checkingAccount, amount)
			if got {
				t.Errorf("PayWithCheck = %t; want false", got)
			}
		})

}
