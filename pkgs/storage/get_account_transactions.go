package storage

import (
	"errors"
)

func (i *InMemoryStorage) GetAccountTransactions(id AccountID) ([]BankTransaction, error) {
	// Prevent NPE: ensure account exists.
	account, accountExists := i.Accounts[id]
	if !accountExists {
		return nil, errors.New("account doesn't exist")
	}

	return account.Transactions, nil
}
