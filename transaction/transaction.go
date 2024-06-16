package transaction

import (
	"errors"
	"github.com/tredoc/go-cor/balance"
	"github.com/tredoc/go-cor/merchant"
	"github.com/tredoc/go-cor/user"
	"sort"
	"sync"
)

type Transaction struct {
	ID       int
	User     *user.User
	Merchant *merchant.Merchant
	Amount   float64
}

func NewTransaction(ID int, u *user.User, m *merchant.Merchant, amount float64) (*Transaction, error) {
	if ID == 0 {
		return nil, errors.New("ID cannot be empty")
	}

	if u == nil {
		return nil, errors.New("user cannot be empty")
	}

	if m == nil {
		return nil, errors.New("merchant cannot be empty")
	}

	if amount == 0 {
		return nil, errors.New("amount cannot be zero")
	}

	return &Transaction{ID: ID, User: u, Merchant: m, Amount: amount}, nil
}

func (t *Transaction) Transfer() error {
	if t.User.Balance.ID == t.Merchant.Balance.ID {
		return errors.New("cannot transfer to the same account")
	}

	type operation struct {
		balance *balance.Balance
		amount  float64
	}

	operations := []operation{
		{balance: t.User.Balance, amount: -t.Amount},
		{balance: t.Merchant.Balance, amount: t.Amount},
	}

	sort.Slice(operations, func(i, j int) bool {
		return operations[i].balance.ID < operations[j].balance.ID
	})

	var wg sync.WaitGroup
	wg.Add(2)

	errChan := make(chan error, 1)

	for _, op := range operations {
		go func() {
			defer wg.Done()
			op.balance.Mu.Lock()
			defer op.balance.Mu.Unlock()

			if op.balance.Amount+op.amount < 0 {
				errChan <- errors.New("insufficient balance")
				return
			}

			op.balance.Amount += op.amount
		}()
	}

	wg.Wait()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}
