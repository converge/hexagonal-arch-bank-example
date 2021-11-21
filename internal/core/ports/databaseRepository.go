package ports

import (
	"hexagonal-example/internal/core/domain/bank"
)

type DatabaseRepository interface {
	Get(accountId int) (*bank.Account, error)
	Save(account *bank.Account) error
}