package account

import "testing"

func TestRandomAccount(t *testing.T) {
	a := Acccount{}
	address := a.NewAddress()
	if address == "" {
		t.Error("NewAddress() should return a non-empty string")
	}
}
