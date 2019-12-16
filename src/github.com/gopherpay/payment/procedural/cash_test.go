//go test github.com/gopherpay/payment/procedural
package procedural

import "testing"

func TestPayWithCash(t *testing.T) {
	got := PayWithCash(100)
	if !got {
		t.Errorf("PayWithCash = %t; want true", got)
	}
}
