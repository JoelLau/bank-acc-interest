package storage

import (
	"time"
)

type Storage interface {
	InsertBankTransaction(InsertBankTransactionParams) (BankTransaction, error)
	GetAccountTransactions(AccountID) ([]BankTransaction, error)

	UpsertInterestRule(UpsertInterestRuleParams) (InterestRule, error)
	GetInterestRules() ([]InterestRule, error)

	GetAccountStatementByMonth(AccountID, time.Time) ([]BankTransaction, error)
}
