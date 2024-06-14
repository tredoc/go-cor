package payment

import (
	"github.com/tredoc/go-cor/bonus"
	"github.com/tredoc/go-cor/notify"
	"github.com/tredoc/go-cor/transaction"
)

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(t *transaction.Transaction)
}

type TransferHandler struct {
	next Handler
}

func (t *TransferHandler) SetNext(handler Handler) Handler {
	t.next = handler
	return handler
}

func (t *TransferHandler) Handle(tr *transaction.Transaction) error {
	err := tr.Handle()
	if err != nil {
		return err
	}

	if t.next == nil {
		t.next.Handle(tr)
	}

	return nil
}

type BonusHandler struct {
	next Handler
}

func (b *BonusHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BonusHandler) Handle(tr *transaction.Transaction) error {
	bns, err := bonus.NewBonus(tr)
	if err != nil {
		return err
	}

	err = bns.Handle()
	if err != nil {
		return err
	}

	if b.next == nil {
		b.next.Handle(tr)
	}

	return nil
}

type NotifyHandler struct {
	next Handler
}

func (n *NotifyHandler) SetNext(handler Handler) Handler {
	n.next = handler
	return handler
}

func (n *NotifyHandler) Handle(tr *transaction.Transaction) error {
	nf, err := notify.NewNotify(tr)
	if err != nil {
		return err
	}

	err = nf.Handle()
	if err != nil {
		return err
	}

	if n.next == nil {
		n.next.Handle(tr)
	}

	return nil
}
