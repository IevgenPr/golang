package oop

import "testing"

func TestCreateCheckingAccount(t *testing.T) {
	want := &CheckingAccount{"Name", "123", "234", 500}
	got := CreateCheckingAccount("Name", "123", "234")
	if got != want {
		t.Errorf("CreateCheckingAccount = %s; want %s", got, want)
	}
}
