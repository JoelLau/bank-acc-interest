package repository

import (
	interestrule "bank-acc-interest/pkgs/interest-rule"
	"bank-acc-interest/pkgs/transactions"
	"errors"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

type Repository interface {
	// append-only.
	// persists transactions.
	//
	// throws errors if balance drops below 0.
	CreateTransaction(transactions.Transaction) (transactions.TransactionID, error)

	// fetch transactions by account id.
	FetchTransactionsByAccountID(transactions.AccountID) ([]transactions.Transaction, error)

	// overrides rules that exist on the same day.
	CreateOrUpdateInterestRule(interestrule.InterestRule) (interestrule.InterestRuleID, error)
}
