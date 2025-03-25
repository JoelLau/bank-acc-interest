package repository

import (
	interestrule "bank-acc-interest/pkgs/interest-rule"
	"bank-acc-interest/pkgs/transactions"
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

type InMemoryRepository struct {
	// NOTE: there's probably a better data structure than a slice
	Transactions []transactions.Transaction
}

func (r *InMemoryRepository) CreateTransaction(transactions.Transaction) (transactions.TransactionID, error) {
	return "", nil
}

func (r *InMemoryRepository) FetchTransactionsByAccountID(transactions.AccountID) ([]transactions.Transaction, error) {
	return nil, nil
}

func (r *InMemoryRepository) CreateOrUpdateInterestRule(interestrule.InterestRule) (interestrule.InterestRuleID, error) {
	return "", nil
}
