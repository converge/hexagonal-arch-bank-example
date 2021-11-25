package ports

import (
	"github.com/google/uuid"
	"hexagonal-example/internal/core/domain/bank"
)

type BankServiceInterface interface {
	Balance(id uuid.UUID) (float64, error)
	Create(bank.Account) (uuid.UUID, error)
}