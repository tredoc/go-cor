package balance

import (
	"errors"
	"sync"
)

type Balance struct {
	ID     int
	UserID int
	Amount float64
	Mu     sync.Mutex
}

func NewBalance(ID int, userID int, amount float64) (*Balance, error) {
	if ID == 0 {
		return nil, errors.New("ID cannot be empty")
	}

	if userID == 0 {
		return nil, errors.New("userID cannot be empty")
	}

	if amount == 0 {
		return nil, errors.New("amount cannot be empty")
	}

	return &Balance{ID: ID, UserID: userID, Amount: amount}, nil
}
