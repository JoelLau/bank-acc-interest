package storage

import (
	"time"

	"github.com/shopspring/decimal"
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
