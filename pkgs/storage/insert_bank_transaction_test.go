package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestInsertBankTransaction(t *testing.T) {
	t.Parallel()

	t.Run("Example", func(t *testing.T) {
		store := storage.NewInMemoryStorage()

		given := storage.InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      storage.TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		}

		have, err := store.InsertBankTransaction(given)
		require.NoError(t, err)

		account, ok := store.Accounts[given.AccountID]
		require.NotNil(t, account)
		require.True(t, ok)
		require.Len(t, store.Accounts[given.AccountID].Transactions, 1)

		require.Equal(t, "20230626-01", have.ID)
		require.Equal(t, given.Type, have.Type)
		require.True(t, decimal.NewFromInt(100).Equal(have.Amount))
		require.True(t, given.Date.Equal(have.Date))

		have, err = store.InsertBankTransaction(given)
		require.NoError(t, err)

		account, ok = store.Accounts[given.AccountID]
		require.NotNil(t, account)
		require.True(t, ok)
		require.Len(t, store.Accounts[given.AccountID].Transactions, 2)

		require.Equal(t, "20230626-02", have.ID)
		require.Equal(t, given.Type, have.Type)
		require.True(t, decimal.NewFromInt(100).Equal(have.Amount))
		require.True(t, given.Date.Equal(have.Date))

	})

	t.Run("2 decimal places", func(t *testing.T) {
		store := storage.NewInMemoryStorage()

		// 20230626 AC001 D 100.001
		given := storage.InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      storage.TransactionTypeDeposit,
			Amount:    decimal.NewFromFloat(100.001),
		}

		_, err := store.InsertBankTransaction(given)
		require.Error(t, err)
	})

	t.Run("input negative amount", func(t *testing.T) {
		store := storage.NewInMemoryStorage()

		// 20230626 AC001 W -100.00
		given := storage.InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      storage.TransactionTypeWithdraw,
			Amount:    decimal.NewFromInt(-100),
		}

		_, err := store.InsertBankTransaction(given)
		require.Error(t, err)
	})

	t.Run("overall negative balance", func(t *testing.T) {
		store := storage.NewInMemoryStorage()

		_, err := store.InsertBankTransaction(storage.InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      storage.TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		})
		require.NoError(t, err)

		_, err = store.InsertBankTransaction(storage.InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 25, 0, 0, 0, 0, time.UTC),
			Type:      storage.TransactionTypeWithdraw,
			Amount:    decimal.NewFromInt(50),
		})
		require.Error(t, err)
	})
}
