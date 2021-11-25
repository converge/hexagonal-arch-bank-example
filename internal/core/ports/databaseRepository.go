package ports

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/core/domain/bank"
)

type DatabaseRepository interface {
	Get(accountId uuid.UUID) (*bank.Account, error)
	Save(account *bank.Account) error
}