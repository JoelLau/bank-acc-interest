package storage

import (
	"time"

	"github.com/shopspring/decimal"
)

type Storage interface {
	InsertBankTransaction(InsertBankTransactionParams) (BankTransaction, error)
	GetAccountTransactionsByMonth(AccountID, time.Time) ([]BankTransaction, error)
	InsertInterestRule(InsertInterestRuleParams) (InterestRule, error)
}

type InsertBankTransactionParams struct {
	AccountID AccountID
	Date      time.Time // "day" level precision
	Type      TransactionType
	Amount    decimal.Decimal
}

type InsertInterestRuleParams struct {
	Date         time.Time       // "date" level precision
	RuleID       string          // user defined
	InterestRate decimal.Decimal // precise up to 2DP
}
