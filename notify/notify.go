package notify

import (
	"errors"
	"github.com/tredoc/go-cor/merchant"
	"github.com/tredoc/go-cor/transaction"
)

type Notify struct {
	Transaction *transaction.Transaction
}

func NewNotify(t *transaction.Transaction) (*Notify, error) {
	if t == nil {
		return nil, errors.New("transfer cannot be empty")
	}

	return &Notify{Transaction: t}, nil
}

func (n *Notify) Handle() error {
	t := n.Transaction.Merchant.NotificationType

	switch t {
	case merchant.Email:
		// send email
		return nil
	case merchant.SMS:
		// send sms
		return nil
	default:
		return errors.New("unknown notification type")
	}
}
