package storage

import (
	"time"

	"github.com/shopspring/decimal"
)

type InterestRule struct {
	Date         time.Time       // "date" level precision
	RuleID       string          // user defined
	InterestRate decimal.Decimal // precise up to 2DP
}
