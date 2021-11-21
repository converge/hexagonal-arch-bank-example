package service

import (
	"hexagonal-example/internal/core/domain/bank"
	"hexagonal-example/internal/repositories"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	srvInstance := &service{}

	if reflect.TypeOf(srv) != reflect.TypeOf(srvInstance) {
		t.Errorf("wanted: instance of *service, got: instance of %v", reflect.TypeOf(srvInstance))
	}
}

func TestServiceWithdrawFromAccount(t *testing.T) {

	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	acc := bank.Account{
		Id: 1,
		Money: 100,
	}

	err := srv.databaseRepository.Save(&acc)
	if err != nil {
		t.Error(err)
	}

	err = srv.WithdrawFromAccount(1, 100)
	if err != nil {
		t.Error(err)
	}

	if acc.Balance() != 0 {
		t.Error("amount should be zero")
	}
}

func TestServiceBalance(t *testing.T) {

	memDb := repositories.NewMemoryDb()
	srv := New(memDb)

	wanted := 100.1

	acc := bank.Account{
		Id: 1,
		Money: wanted,
	}

	err := srv.databaseRepository.Save(&acc)
	if err != nil {
		t.Error(err)
	}

	balance, err := srv.Balance(1)
	if err != nil {
		t.Error(err)
	}

	if wanted != balance {
		t.Errorf("wanted: %T, got: %T", wanted, balance)
	}
}