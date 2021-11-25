package services

import (
	"github.com/google/uuid"
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

func (srv service) WithdrawFromAccount(id uuid.UUID, amount float64) error {
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

func (srv service) Balance(id uuid.UUID) (float64, error) {
	account, err := srv.databaseRepository.Get(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return account.Balance(), nil
}

func (srv service) Create(account bank.Account) (uuid.UUID, error) {
	account.Id = uuid.New()
	err := srv.databaseRepository.Save(&account)
	if err != nil {
		log.Println(err)
		return uuid.Nil, err
	}
	return account.Id, nil
}