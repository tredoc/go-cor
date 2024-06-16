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

	t3, err := utils.GetTransaction(555)
	if err != nil {
		test.Fatal(err)
	}

	tests := []struct {
		name string
		t    *transaction.Transaction
		bb   float64
	}{
		{
			name: "Test accrue 10 bonus",
			t:    t1,
			bb:   t1.User.BonusBalance.Amount,
		},
		{
			name: "Test not accrue bonus",
			t:    t2,
			bb:   t1.User.BonusBalance.Amount,
		},
		{
			name: "Test accrue 10% bonus",
			t:    t3,
			bb:   t1.User.BonusBalance.Amount,
		},
	}

	for _, tt := range tests {
		test.Run(tt.name, func(t *testing.T) {
			threshold := 100.0
			percent := 10.0
			b, err := bonus.NewBonus(tt.t, threshold, percent)
			if err != nil {
				t.Fatal(err)
			}

			b.Accrue()
			if tt.t.Amount > threshold && utils.AlmostEqual(tt.t.User.BonusBalance.Amount, tt.bb+tt.t.Amount, percent/10000000.0) {
				t.Errorf("user bonus balance is %f, want %f", tt.t.User.BonusBalance.Amount, tt.bb+tt.t.Amount*(percent/100.0))
			}
		})
	}
}
