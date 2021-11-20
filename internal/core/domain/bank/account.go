package bank

// internal/core/domain/bank

type Account struct {
	Id    int
	Money float64
}

//func NewAccount(id int, initialMoney float64) *Account {
//	return &Account{
//		Id:    id,
//		Money: initialMoney,
//	}
//}

func (acc *Account) Balance() (balance float64) {
	return acc.Money
}

func (acc *Account) Withdraw(amount float64) error {
	acc.Money = acc.Balance() - amount

	return nil
}
