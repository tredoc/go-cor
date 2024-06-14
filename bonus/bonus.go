package bonus

import (
	"errors"
	"github.com/tredoc/go-cor/transaction"
)

type Bonus struct {
	Transaction *transaction.Transaction
}

func NewBonus(t *transaction.Transaction) (*Bonus, error) {

	if t == nil {
		return nil, errors.New("transfer cannot be empty")
	}

	return &Bonus{Transaction: t}, nil
}

func (b *Bonus) Handle() error {
	if b.Transaction.Amount < 100 {
		return errors.New("transfer amount must be greater than 100")
	}

	b.Transaction.User.BonusBalance.Mu.Lock()
	defer b.Transaction.User.BonusBalance.Mu.Unlock()

	b.Transaction.User.BonusBalance.Amount += 10

	return nil
}