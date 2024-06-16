package merchant

import (
	"errors"
	"github.com/tredoc/go-cor/balance"
)

type Merchant struct {
	ID      int
	Name    string
	UserID  int
	Balance *balance.Balance
}

func NewMerchant(ID int, name string, userID int, balance *balance.Balance) (*Merchant, error) {
	if ID == 0 {
		return nil, errors.New("ID cannot be empty")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if userID == 0 {
		return nil, errors.New("userID cannot be empty")
	}

	if balance == nil {
		return nil, errors.New("balance cannot be empty")
	}

	return &Merchant{ID: ID, Name: name, UserID: userID, Balance: balance}, nil
}
