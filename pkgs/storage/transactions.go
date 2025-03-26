package storage

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	TransactionID = string
	AccountID     = string
)

type Account struct {
	ID           AccountID
	Transactions []BankTransaction // must main sortedness
}

type BankTransaction struct {
	// Populated by DB.
	//
	// In YYYYMMdd-xx format, where xx is a running number.
	ID TransactionID

	Type TransactionType

	// precision is only up to "day" level
	Date time.Time

	Amount decimal.Decimal

	Balance decimal.Decimal
}

type TransactionType string

const (
	TransactionTypeWithdraw TransactionType = "W"
	TransactionTypeDeposit  TransactionType = "D"
	TransactionTypeInterest TransactionType = "I"
)
