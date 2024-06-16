package payment

import (
	"github.com/tredoc/go-cor/bonus"
	"github.com/tredoc/go-cor/notification"
	"github.com/tredoc/go-cor/transaction"
)

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(t *transaction.Transaction) error
}

type TransferHandler struct {
	next Handler
}

func (th *TransferHandler) SetNext(handler Handler) Handler {
	th.next = handler
	return handler
}

func (th *TransferHandler) Handle(t *transaction.Transaction) error {
	err := t.Transfer()
	if err != nil {
		return err
	}

	if th.next == nil {
		return nil
	}
	err = th.next.Handle(t)
	if err != nil {
		return err
	}

	return nil
}

type BonusHandler struct {
	next Handler
}

func (bh *BonusHandler) SetNext(handler Handler) Handler {
	bh.next = handler
	return handler
}

func (bh *BonusHandler) Handle(t *transaction.Transaction) error {
	bns, err := bonus.NewBonus(t, 100.0, 10.0)
	if err != nil {
		return err
	}

	bns.Accrue()

	if bh.next == nil {
		return nil
	}
	err = bh.next.Handle(t)
	if err != nil {
		return err
	}

	return nil
}

type NotificationHandler struct {
	next Handler
}

func (nh *NotificationHandler) SetNext(handler Handler) Handler {
	nh.next = handler
	return handler
}

func (nh *NotificationHandler) Handle(t *transaction.Transaction) error {
	nf, err := notification.NewNotification(t)
	if err != nil {
		return err
	}

	err = nf.Send()
	if err != nil {
		return err
	}

	if nh.next == nil {
		return nil
	}
	err = nh.next.Handle(t)
	if err != nil {
		return err
	}

	return nil
}
