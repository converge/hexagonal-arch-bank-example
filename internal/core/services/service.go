package services

import (
	"fmt"
	"hexagonal-example/internal/core/domain/bank"
	"hexagonal-example/internal/core/ports"
	"log"
)

type service struct {
	databaseRepository ports.DatabaseRepository
}

func New(databaseRepository ports.DatabaseRepository) *service {
	return &service{
		databaseRepository: databaseRepository,
	}
}

func (srv service) WithdrawFromAccount(id int, amount float64) error {
	account, err := srv.databaseRepository.Get(id)
	if err != nil {
		return err
	}

	err = account.Withdraw(amount)
	if err != nil {
		return err
	}

	return nil
}

func (srv service) Balance(id int) (float64, error) {
	account, err := srv.databaseRepository.Get(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return account.Balance(), nil
}

func (srv service) Create(id int, money float64) bank.Account {
	fmt.Println("hiii")
	return bank.Account{
		Id: 1,
		Money: 100,
	}
}