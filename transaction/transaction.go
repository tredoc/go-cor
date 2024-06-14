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
		return nil, errors.New("amount cannot be empty")
	}

	return &Transaction{ID: ID, User: u, Merchant: m, Amount: amount}, nil
}

func (t *Transaction) Handle() error {
	if t.User.Balance.ID == t.Merchant.Balance.ID {
		return errors.New("cannot transfer to the same account")
	}

	type Operation struct {
		Balance *balance.Balance
		Amount  float64
	}

	operations := []Operation{
		{Balance: t.User.Balance, Amount: -t.Amount},
		{Balance: t.Merchant.Balance, Amount: t.Amount},
	}

	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Balance.ID < operations[j].Balance.ID
	})

	var wg sync.WaitGroup
	wg.Add(2)

	errChan := make(chan error, 1)

	for _, op := range operations {
		go func() {
			op.Balance.Mu.Lock()
			defer op.Balance.Mu.Unlock()
			if op.Balance.Amount+op.Amount < 0 {
				errChan <- errors.New("insufficient balance")
				return
			}
		}()
	}

	wg.Wait()

	err := <-errChan
	if err != nil {
		return err
	}

	return nil
}
