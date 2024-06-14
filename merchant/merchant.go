package merchant

import (
	"errors"
	"github.com/tredoc/go-cor/balance"
)

type Notification string

const (
	Email Notification = "email"
	SMS   Notification = "sms"
)

type Merchant struct {
	ID               int
	Name             string
	UserID           int
	Balance          *balance.Balance
	NotificationType Notification
}

func NewMerchant(ID int, name string, userID int, balance *balance.Balance, notificationType Notification) (*Merchant, error) {
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

	if notificationType == "" {
		return nil, errors.New("notificationType cannot be empty")
	}

	return &Merchant{ID: ID, Name: name, UserID: userID, Balance: balance, NotificationType: notificationType}, nil
}
