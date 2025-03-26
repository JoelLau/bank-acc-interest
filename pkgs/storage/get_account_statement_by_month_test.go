package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAccountStatementByMonth(t *testing.T) {
	t.Parallel()

	t.Run("Example from REQUIREMENT.md", func(t *testing.T) {
		t.Parallel()

		accountID := "AC001"
		store := storage.NewInMemoryStorage()

		// Starting balance of 100 from May 5th deposit
		store.Accounts[accountID] = storage.Account{
			ID: accountID,
			Transactions: []storage.BankTransaction{
				{
					Date:   time.Date(2023, 05, 05, 0, 0, 0, 0, time.UTC),
					ID:     "20230505-01",
					Type:   storage.TransactionTypeDeposit,
					Amount: decimal.NewFromFloat(100.00),
				},
				{
					Date:   time.Date(2023, 06, 01, 0, 0, 0, 0, time.UTC),
					ID:     "20230601-01",
					Type:   storage.TransactionTypeDeposit,
					Amount: decimal.NewFromInt(150),
				},
				{
					Date:   time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
					ID:     "20230626-01",
					Type:   storage.TransactionTypeWithdraw,
					Amount: decimal.NewFromInt(20),
				},
				{
					Date:   time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
					ID:     "20230626-02",
					Type:   storage.TransactionTypeWithdraw,
					Amount: decimal.NewFromInt(100),
				},
			},
		}

		store.InterestRules = []storage.InterestRule{
			{
				Date:         time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				RuleID:       "RULE01",
				InterestRate: decimal.NewFromFloat(1.95),
			},
			{
				Date:         time.Date(2023, 05, 20, 0, 0, 0, 0, time.UTC),
				RuleID:       "RULE02",
				InterestRate: decimal.NewFromFloat(1.90),
			},
			{
				Date:         time.Date(2023, 06, 15, 0, 0, 0, 0, time.UTC),
				RuleID:       "RULE03",
				InterestRate: decimal.NewFromFloat(2.20),
			},
		}

		statements, err := store.GetAccountStatementByMonth(accountID, time.Date(2023, 06, 01, 0, 0, 0, 0, time.UTC))
		require.NoError(t, err)
		require.GreaterOrEqual(t, 4, len(statements))

		interestTxn := statements[len(statements)-1]
		assert.Equal(t, storage.TransactionTypeInterest, interestTxn.Type)
		assert.Equal(t, "0.39", interestTxn.Amount.StringFixed(2))
		assert.Equal(t, "130.39", interestTxn.Balance.StringFixed(2))
	})
}
