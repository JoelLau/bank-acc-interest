package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetAccountTransactions(t *testing.T) {
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
		}

		have, err := store.GetAccountTransactions(accountID)
		require.NoError(t, err)
		require.Equal(t, have, want)
	})

	t.Run("no account", func(t *testing.T) {
		store := storage.NewInMemoryStorage()
		_, err := store.GetAccountTransactions("non-existent account")
		require.Error(t, err)
	})
}
