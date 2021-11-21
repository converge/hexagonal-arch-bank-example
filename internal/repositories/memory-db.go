package repositories

import (
	"errors"
	"hexagonal-example/internal/core/domain/bank"
)

type memoryDb struct {
	instance map[int]*bank.Account
}

func NewMemoryDb() *memoryDb {
	memDb := make(map[int]*bank.Account)
	return &memoryDb{
		instance: memDb,
	}
}

func (memDb *memoryDb) Get(accountId int) (*bank.Account, error) {

	account := memDb.instance[accountId]

	if account == nil {
		return nil, errors.New("account doesnt exist")
	}

	return account, nil
}

func (memDb *memoryDb) Save(account *bank.Account) error {
	memDb.instance[account.Id] = account
	return nil
}