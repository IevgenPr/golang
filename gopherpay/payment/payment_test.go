package payment

//go test github.com/gopherpay/payment/procedural -v
import "testing"

func TestPayWithCredit(t *testing.T) {
	// set up
	const name, cardNum = "Name Boring", "1111-2222-3333-4444"
	const month, year, securityCode = 8, 12, 123

	credit := CreateCreditAccount(name, cardNum, month, year, securityCode)

	t.Run("Credit owner name is Name Boring",
		func(t *testing.T) {
			if credit.OwnerName() != name {
				t.Errorf("OwnerName = %s; want %s", credit.OwnerName(), name)
			}
		})
	t.Run("Card number is correct",
		func(t *testing.T) {
			if credit.CardNumber() != cardNum {
				t.Errorf("CardNumber = %s; want %s", credit.CardNumber(), cardNum)
			}
		})
	t.Run("Change to invalid Card number",
		func(t *testing.T) {
			if credit.SetCardNumber("invalid").Error() != "Invalid card number format" {
				t.Errorf("Invalid card number accepted")
			}
		})
	t.Run("Check available credit",
		func(t *testing.T) {
			if credit.AvailableCredit() != 5000. {
				t.Errorf("AvailableCredit = %f, expected 5000.", credit.AvailableCredit())
			}
		})
	// got := PayWithCredit(credit, 100)
	// if !got {
	// 	t.Errorf("PayWithCash = %t; want true", got)
	// }

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
