package storage

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func TestGetLatestBankTransactionByAccount(t *testing.T) {
	t.Parallel()

	t.Run("insert to empty database", func(t *testing.T) {
		t.Parallel()

		s := NewInMemoryStorage()
		require.Len(t, s.BankTransactions, 0)

		have, err := s.InsertBankTransaction(InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		})
		require.NoError(t, err)

		require.Equalf(t, "20230626-01", have.ID, "ID mismatch")
		require.Equalf(t, "AC001", have.AccountID, "AccountID mismatch")
		require.Equalf(t, time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC), have.Date, "Date mismatch")
		require.Equalf(t, TransactionTypeDeposit, have.Type, "Date mismatch")
		require.False(t, have.CreatedAt.IsZero(), "CreatedAt should be populated")

		require.Len(t, s.BankTransactions, 1)
	})

	t.Run("increment id", func(t *testing.T) {
		t.Parallel()

		s := NewInMemoryStorage()
		require.Len(t, s.BankTransactions, 0)

		have, err := s.InsertBankTransaction(InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		})
		require.NoError(t, err)
		require.Equalf(t, "20230626-01", have.ID, "ID mismatch")
		require.Len(t, s.BankTransactions, 1)

		have, err = s.InsertBankTransaction(InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 26, 0, 0, 0, 0, time.UTC),
			Type:      TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		})
		require.NoError(t, err)

		require.Equalf(t, "20230626-02", have.ID, "ID mismatch")
		require.Len(t, s.BankTransactions, 2)

		have, err = s.InsertBankTransaction(InsertBankTransactionParams{
			AccountID: "AC001",
			Date:      time.Date(2023, 06, 27, 0, 0, 0, 0, time.UTC),
			Type:      TransactionTypeDeposit,
			Amount:    decimal.NewFromInt(100),
		})
		require.NoError(t, err)

		require.Equalf(t, "20230627-01", have.ID, "ID mismatch")
		require.Len(t, s.BankTransactions, 3)
	})
}
