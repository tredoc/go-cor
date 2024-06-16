package utils

import (
	"errors"
	"github.com/tredoc/go-cor/balance"
	"github.com/tredoc/go-cor/merchant"
	"github.com/tredoc/go-cor/transaction"
	"github.com/tredoc/go-cor/user"
)

func GetUser(amount float64) (*user.User, error) {
	bal, err := balance.NewBalance(1, 1, amount)
	if err != nil {
		return nil, err
	}

	bonusBal, err := balance.NewBalance(2, 1, 0)
	if err != nil {
		return nil, err
	}

	u, err := user.NewUser(1, "password", "username", bal, bonusBal)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetMerchant(amount float64) (*merchant.Merchant, error) {
	merchantBal, err := balance.NewBalance(3, 3, amount)
	if err != nil {
		return nil, err
	}

	m, err := merchant.NewMerchant(2, "prof pay", 2, merchantBal)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func GetTransaction(amount float64) (*transaction.Transaction, error) {
	ub, err := balance.NewBalance(1, 1, 100.0)
	if err != nil {
		return nil, err
	}

	ubb, err := balance.NewBalance(1, 1, 100.0)
	if err != nil {
		return nil, err
	}
	u, err := user.NewUser(1, "password", "username", ub, ubb)
	if err != nil {
		return nil, err
	}

	mb, err := balance.NewBalance(5, 2, 100.0)
	if err != nil {
		return nil, err
	}

	m, err := merchant.NewMerchant(1, "merchant", 2, mb)

	t, err := transaction.NewTransaction(1, u, m, amount)
	if err != nil {
		return nil, err
	}

	if t == nil {
		return nil, errors.New("transaction is nil")
	}

	return t, nil
}
