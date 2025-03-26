package storage

import (
	"time"

	decimal "github.com/jackc/pgx-shopspring-decimal"
)

type Storage interface {
	InsertBankTransaction(InsertBankTransactionParams) (BankTransaction, error)
}

type InsertBankTransactionParams struct {
	AccountID AccountID
	Date      time.Time // "day" level precision
	Type      TransactionType
	Amount    decimal.Decimal
}
