package ports

import (
	"hexagonal-example/internal/core/domain/bank"
)

type DatabaseRepository interface {
	// is that ok that bank.Account is exported?
	Get(accountId int) (*bank.Account, error)
	Save(account *bank.Account) error
}