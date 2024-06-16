package bonus_test

import (
	"github.com/tredoc/go-cor/bonus"
	"github.com/tredoc/go-cor/transaction"
	"github.com/tredoc/go-cor/utils"
	"testing"
)

func TestBonus_Accrue(test *testing.T) {
	test.Parallel()

	t1, err := utils.GetTransaction(100)
	if err != nil {
		test.Fatal(err)
	}

	t2, err := utils.GetTransaction(99)
	if err != nil {
		test.Fatal(err)
	}

	bb := t1.User.BonusBalance.Amount

	tests := []struct {
		name string
		t    *transaction.Transaction
	}{
		{
			name: "Test accrue bonus",
			t:    t1,
		},
		{
			name: "Test not accrue bonus",
			t:    t2,
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			b, err := bonus.NewBonus(tt.t)
			if err != nil {
				t.Fatal(err)
			}

			b.Accrue()
			if tt.t.Amount > 100 && tt.t.User.BonusBalance.Amount != bb+tt.t.Amount {
				t.Errorf("Bonus.Accrue() = %v, want %v", tt.t.User.BonusBalance.Amount, tt.t.Amount)
			}
		})
	}
}
