package ports

import "hexagonal-example/internal/core/domain/bank"

type BankServiceInterface interface {
	Balance(id int) (float64, error)
	Create(id int, money float64) bank.Account
}