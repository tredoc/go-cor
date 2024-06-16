package payment_test

import (
	"github.com/tredoc/go-cor/payment"
	"github.com/tredoc/go-cor/utils"
	"testing"
)

func TestPayment(test *testing.T) {
	th := payment.TransferHandler{}
	bh := payment.BonusHandler{}
	nh := payment.NotificationHandler{}

	th.SetNext(&bh).SetNext(&nh)

	test.Run("Test payment", func(test *testing.T) {
		bonusAmount := 10.0
		amount := 100.0
		t, err := utils.GetTransaction(amount)
		if err != nil {
			test.Fatal(err)
		}

		ub := t.User.Balance.Amount
		ubb := t.User.BonusBalance.Amount
		mb := t.Merchant.Balance.Amount

		err = th.Handle(t)
		if err != nil {
			test.Fatal(err)
		}

		if t.User.Balance.Amount != ub-amount {
			test.Errorf("expected user balance to be %f, got %v", ub-amount, t.User.Balance.Amount)
		}

		if t.User.BonusBalance.Amount != ubb+bonusAmount {
			test.Errorf("expected user bonus balance to be %f, got %v", ubb+bonusAmount, t.User.BonusBalance.Amount)
		}

		if t.Merchant.Balance.Amount != mb+amount {
			test.Errorf("expected merchant balance to be %f, got %v", mb+amount, t.Merchant.Balance.Amount)
		}
	})
}
