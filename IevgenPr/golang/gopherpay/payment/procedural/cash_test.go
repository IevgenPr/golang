package procedural

//go test github.com/gopherpay/payment/procedural -v

import "testing"

func TestPayWithCash(t *testing.T) {
	got := PayWithCash(100)
	if !got {
		t.Errorf("PayWithCash = %t; want true", got)
	}
}
