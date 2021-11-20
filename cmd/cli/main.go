package main

import (
	"fmt"
	"hexagonal-example/internal/core/domain/bank"
	"hexagonal-example/internal/core/service"
	"hexagonal-example/internal/repositories"
	"log"
)

func main() {
	dbRepository := repositories.NewMemoryDb()

	// todo: should Account.New() create it?
	acc := bank.Account{
		Id: 1,
		Money: 100,
	}

	err := dbRepository.Save(&acc)
	if err != nil {
		log.Println(err)
		return
	}

	srv := service.New(dbRepository)
	err = srv.WithdrawFromAccount(1, 50)
	if err != nil {
		log.Println(err)
		return
	}

	balance, err := srv.Balance(acc.Id)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(balance)

}