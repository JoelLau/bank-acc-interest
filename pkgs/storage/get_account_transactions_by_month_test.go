package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetAccountTransactionsByMonth(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		store := storage.NewInMemoryStorage()
		accountID := "AC001"

		want := []storage.BankTransaction{{
			Date:   time.Date(2023, time.June, 26, 0, 0, 0, 0, time.UTC),
			ID:     "20230606-01",
			Type:   storage.TransactionTypeDeposit,
			Amount: decimal.NewFromInt(100),
		}}

		store.Accounts[accountID] = storage.Account{
			ID:           accountID,
			Transactions: want,
			Balance:      decimal.NewFromInt(100),
		}

		have, err := store.GetAccountTransactionsByMonth(accountID, time.Date(2023, 06, 25, 0, 0, 0, 0, time.UTC))
		require.NoError(t, err)
		require.Equal(t, have, want)
	})

	t.Run("no account", func(t *testing.T) {
		store := storage.NewInMemoryStorage()
		_, err := store.GetAccountTransactionsByMonth("non-existent account", time.Now())
		require.Error(t, err)
	})
}
