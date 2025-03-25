package transactions

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	TransactionID = string
	AccountID     = string
)

type Transaction struct {
	// populated by DB
	ID TransactionID

	AccountID AccountID

	Type TransactionType

	// precision is only up to "day" level
	Date time.Time

	Amount decimal.Decimal
}
