package procedural

//go test github.com/gopherpay/payment/procedural -v
import "testing"

func TestPayWithCredit(t *testing.T) {
	// set up
	const amount = 500
	credit := &CreditCard{
		OwnerName:       "Debra Ingram",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 8,
		ExpirationYear:  2024,
		SecurityCode:    123,
		AvailableCredit: 5000,
	}
	got := PayWithCredit(credit, 100)
	if !got {
		t.Errorf("PayWithCash = %t; want true", got)
	}

	t.Run("Credit 100 from 5000, get 4900 in available credit",
		func(t *testing.T) {
			const remainingCredit = 4400
			got := PayWithCredit(credit, amount)
			if !got {
				t.Errorf("PayWithCheck = %t; want true", got)
			}
			if credit.AvailableCredit != remainingCredit {
				t.Errorf("Available credit = %.2f; want %d",
					credit.AvailableCredit, remainingCredit)
			}
		})

	t.Run("Check failed to withdraw",
		func(t *testing.T) {
			got := PayWithCredit(credit, 5000)
			if got {
				t.Errorf("PayWithCredit = %t; want false", got)
			}
			if credit.AvailableCredit != 4400 {
				t.Errorf("PayWithCredit = %f; want 4400", credit.AvailableCredit)
			}
		})
}
