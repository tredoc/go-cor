package user

import (
	"errors"
	"github.com/tredoc/go-cor/balance"
)

type User struct {
	ID           int
	Username     string
	Password     string
	Balance      *balance.Balance
	BonusBalance *balance.Balance
}

func NewUser(ID int, password string, username string, balance *balance.Balance, bonusBalance *balance.Balance) (*User, error) {
	if ID == 0 {
		return nil, errors.New("ID cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if balance == nil {
		return nil, errors.New("balance cannot be empty")
	}

	if bonusBalance == nil {
		return nil, errors.New("bonus balance cannot be empty")
	}

	return &User{ID: ID, Password: password, Username: username, Balance: balance, BonusBalance: bonusBalance}, nil
}
