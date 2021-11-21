package bank

import "testing"

func TestAccountBalance(t *testing.T) {

	acc := Account{
		Id: 1,
		Money: 100,
	}

	balance := acc.Balance()

	if acc.Money != balance {
		t.Errorf("wanted: %T, got: %T", acc.Money, balance)
	}
}

func TestAccountWithdrawSuccess(t *testing.T) {

	wanted := 0

	acc := Account{
		Id: 1,
		Money: 100,
	}

	err := acc.Withdraw(100)
	if err != nil {
		t.Error(err)
	}

	if acc.Money != 0 {
		t.Errorf("wanted: %T, got: %T", wanted, acc.Money)
	}
}

func TestAccountWithdrawNoFunds(t *testing.T) {

	acc := Account{
		Id: 1,
		Money: 100,
	}

	err := acc.Withdraw(200)
	if err.Error() != "no enough funds" {
		t.Error(err)
	}
}
