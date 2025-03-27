package storage

import (
	"errors"
	"time"
)

func (i *InMemoryStrorage) GetAccountTransactionsByMonth(id AccountID, date time.Time) ([]BankTransaction, error) {
	// Prevent NPE: ensure account exists.
	_, accountExists := i.Accounts[id]
	if !accountExists {
		return nil, errors.New("account doesn't exist")
	}

	transactions := make([]BankTransaction, 0)
	for _, transaction := range i.Accounts[id].Transactions {
		if date.Year() == transaction.Date.Year() &&
			date.Month() == transaction.Date.Month() {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}
