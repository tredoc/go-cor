package bonus

import (
	"errors"
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

	return &Bonus{Transaction: t, Threshold: threshold, Percent: percent}, nil
}

func (b *Bonus) Accrue() {
	if b.Transaction.Amount >= b.Threshold {
		b.Transaction.User.BonusBalance.Mu.Lock()
		defer b.Transaction.User.BonusBalance.Mu.Unlock()

		bonusAmount := b.Transaction.Amount * (b.Percent / 100.0)
		b.Transaction.User.BonusBalance.Amount += bonusAmount
	}
}
