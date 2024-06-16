package bonus

import (
	"errors"
	"fmt"
	"github.com/tredoc/go-cor/transaction"
)

type Bonus struct {
	Threshold   float64
	Percent     float64
	Transaction *transaction.Transaction
}

func NewBonus(t *transaction.Transaction, threshold float64, percent float64) (*Bonus, error) {
	if t == nil {
		return nil, errors.New("transfer cannot be empty")
	}

	if threshold < 0 {
		return nil, errors.New("threshold cannot be negative")
	}

	if percent <= 0 {
		return nil, errors.New("percent cannot be zero or negative")
	}

	return &Bonus{Transaction: t}, nil
}

func (b *Bonus) Accrue() {
	if b.Transaction.Amount >= b.Threshold {
		b.Transaction.User.BonusBalance.Mu.Lock()
		defer b.Transaction.User.BonusBalance.Mu.Unlock()
		fmt.Println("HERE")
		b.Transaction.User.BonusBalance.Amount += (b.Transaction.Amount * (b.Percent / 100.0))
	}
}
