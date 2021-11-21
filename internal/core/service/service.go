package service

import (
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