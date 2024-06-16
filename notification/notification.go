package notification

import (
	"errors"
	"github.com/tredoc/go-cor/transaction"
)

type Type string

const (
	Email Type = "email"
	SMS   Type = "sms"
)

type Notification struct {
	Transaction *transaction.Transaction
}

func NewNotification(t *transaction.Transaction) (*Notification, error) {
	if t == nil {
		return nil, errors.New("transfer cannot be empty")
	}

	return &Notification{Transaction: t}, nil
}

func (n *Notification) Send() error {
	return nil
}
