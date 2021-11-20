package repositories

import (
	"errors"
	"fmt"
	"hexagonal-example/internal/core/domain/bank"
)

type memoryDb struct {
	instance map[int]*bank.Account
}

func NewMemoryDb() *memoryDb {
	// todo: make?
	//return &memoryDb{instance: map[int]bank.Account{}}

	y := make(map[int]*bank.Account)
	return &memoryDb{
		instance: y,
	}
}

func (memDb *memoryDb) Get(accountId int) (*bank.Account, error) {

	account := memDb.instance[accountId]
	fmt.Println(account)

	if account == nil {
		return nil, errors.New("account doesnt exist")
	}

	return account, nil
}

func (memDb *memoryDb) Save(account *bank.Account) error {
	memDb.instance[account.Id] = account
	return nil
}