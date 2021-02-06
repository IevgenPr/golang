package oop

import "testing"

func TestCreateCheckingAccount(t *testing.T) {
	want := CheckingAccount{"Name", "123", "234", 500.0}
	got := *CreateCheckingAccount("Name", "123", "234")
	if got != want {
		t.Errorf("CreateCheckingAccount got/want: \n%v \n%v", got, want)
	}
}
