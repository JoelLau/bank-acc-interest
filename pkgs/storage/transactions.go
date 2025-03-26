package storage

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	TransactionID = string
	AccountID     = string
)

type BankTransaction struct {
	// Populated by DB.
	//
	// In YYYYMMdd-xx format, where xx is a running number.
	ID TransactionID

	// Populated by DB.
	CreatedAt time.Time

	AccountID AccountID

	Type TransactionType

	// precision is only up to "day" level
	Date time.Time

	Amount decimal.Decimal
}

type TransactionType string

const (
	TransactionTypeWidthdraw TransactionType = "W"
	TransactionTypeDeposit   TransactionType = "D"
	TransactionTypeInterest  TransactionType = "I"
)
