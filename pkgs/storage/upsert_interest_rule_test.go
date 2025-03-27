package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpsertInterestRule(t *testing.T) {
	t.Parallel()

	t.Run("override by date", func(t *testing.T) {
		t.Parallel()

		first := storage.UpsertInterestRuleParams{
			RuleID:       "FIRST",
			InterestRate: decimal.NewFromFloat(1.1),
			Date:         time.Date(2023, 06, 15, 0, 0, 0, 0, time.Local),
		}

		store := storage.NewInMemoryStorage()
		require.Len(t, store.InterestRules, 0)

		have, err := store.UpsertInterestRule(first)
		require.NoError(t, err)

		assert.Equal(t, first.RuleID, have.RuleID)
		assert.True(t, first.InterestRate.Equal(have.InterestRate))
		assert.True(t, first.Date.Equal(have.Date))

		require.Len(t, store.InterestRules, 1)

		second := storage.UpsertInterestRuleParams{
			RuleID:       "SECOND",
			InterestRate: decimal.NewFromFloat(2.2),
			Date:         time.Date(2023, 06, 15, 0, 0, 0, 0, time.Local),
		}
		have, err = store.UpsertInterestRule(second)
		require.NoError(t, err)
		require.Len(t, store.InterestRules, 1)

		assert.Equal(t, second.RuleID, have.RuleID)
		assert.True(t, second.InterestRate.Equal(have.InterestRate))
		assert.True(t, second.Date.Equal(have.Date))
	})

	t.Run("rate < 0", func(t *testing.T) {
		t.Parallel()

		first := storage.UpsertInterestRuleParams{
			RuleID:       "RULE01",
			InterestRate: decimal.NewFromFloat(-0.01),
			Date:         time.Date(2023, 06, 15, 0, 0, 0, 0, time.Local),
		}

		store := storage.NewInMemoryStorage()
		require.Len(t, store.InterestRules, 0)

		_, err := store.UpsertInterestRule(first)
		require.Error(t, err)
	})

	t.Run("rate > 100", func(t *testing.T) {
		t.Parallel()

		first := storage.UpsertInterestRuleParams{
			RuleID:       "RULE01",
			InterestRate: decimal.NewFromFloat(100.1),
			Date:         time.Date(2023, 06, 15, 0, 0, 0, 0, time.Local),
		}

		store := storage.NewInMemoryStorage()
		require.Len(t, store.InterestRules, 0)

		_, err := store.UpsertInterestRule(first)
		require.Error(t, err)
	})
}
