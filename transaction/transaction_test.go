package transaction_test

import (
	"github.com/tredoc/go-cor/transaction"
	"github.com/tredoc/go-cor/utils"
	"sync"
	"testing"
)

func TestTransaction_Handle(test *testing.T) {
	test.Parallel()

	test.Run("test single transfer handle", func(test *testing.T) {
		amount := 100.0
		t, err := utils.GetTransaction(amount)
		if err != nil {
			test.Fatal("setup returned not valid transaction: ", err.Error())
		}

		ub := t.User.Balance.Amount
		mb := t.Merchant.Balance.Amount

		err = t.Transfer()
		if err != nil {
			test.Errorf("error: %v", err)
		}

		if t.User.Balance.Amount != ub-amount {
			test.Errorf("expected user balance to be %f, got %v", ub-amount, t.User.Balance.Amount)
		}

		if t.Merchant.Balance.Amount != mb+amount {
			test.Errorf("expected merchant balance to be %f, got %v", mb+amount, t.Merchant.Balance.Amount)
		}
	})

	test.Run("test single transfer with insufficient funds", func(test *testing.T) {
		amount := 10000000000.0
		t, err := utils.GetTransaction(amount)
		if err != nil {
			test.Fatal("setup returned not valid transaction: ", err.Error())
		}

		err = t.Transfer()
		if err == nil {
			test.Error("expected error, got nil")
		}

		if err.Error() != "insufficient balance" {
			test.Errorf("expected error to be 'insufficient balance', got %v", err.Error())
		}
	})

	test.Run("test concurrent transfer", func(test *testing.T) {
		userAmount := 1000.0
		merchantAmount := 1000.0

		u, err := utils.GetUser(userAmount)
		if err != nil {
			test.Fatal("setup returned not valid user: ", err.Error())
		}

		m, err := utils.GetMerchant(merchantAmount)
		if err != nil {
			test.Fatal("setup returned not valid merchant: ", err.Error())
		}

		amount := 1.0
		times := 1000

		var wg sync.WaitGroup
		wg.Add(times)

		for i := 1; i <= times; i++ {
			go func() {
				defer wg.Done()
				t, err := transaction.NewTransaction(i, u, m, amount)
				if err != nil {
					test.Errorf("error: %v", err)
				}
				err = t.Transfer()
				if err != nil {
					test.Errorf("error: %v", err)
				}
			}()
		}

		wg.Wait()

		if u.Balance.Amount != userAmount-float64(times)*amount {
			test.Errorf("expected user balance to be %v, got %v", userAmount-float64(times)*amount, u.Balance.Amount)
		}

		if m.Balance.Amount != merchantAmount+float64(times)*amount {
			test.Errorf("expected merchant balance to be %v, got %v", merchantAmount+float64(times)*amount, m.Balance.Amount)
		}
	})
}
