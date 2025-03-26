package storage

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/shopspring/decimal"
)

type InsertBankTransactionParams struct {
	AccountID AccountID
	Date      time.Time // "day" level precision
	Type      TransactionType
	Amount    decimal.Decimal
}

func (i *InMemoryStorage) InsertBankTransaction(params InsertBankTransactionParams) (BankTransaction, error) {
	// Prevent NPE: ensure account exists.
	_, accountExists := i.Accounts[params.AccountID]
	if !accountExists {
		i.Accounts[params.AccountID] = Account{
			ID:           params.AccountID,
			Transactions: make([]BankTransaction, 0),
		}
	}

	// evaluate new balance (i.e. "BEGIN TRANSACTION; ROLLBACK")
	transactions := make([]BankTransaction, len(i.Accounts[params.AccountID].Transactions))
	copy(transactions, i.Accounts[params.AccountID].Transactions)

	newTransaction := BankTransaction{
		ID:     newBankTransactionID(transactions, params.Date),
		Type:   params.Type,
		Date:   params.Date,
		Amount: params.Amount,
	}

	// no negative numbers
	if newTransaction.Amount.IsNegative() {
		return BankTransaction{}, errors.New("invalid input")
	}

	// 2dp precision
	if newTransaction.Amount.Exponent() < -2 {
		return BankTransaction{}, errors.New("invalid input")
	}

	transactions = append(transactions, newTransaction)

	slices.SortFunc(transactions, func(a, b BankTransaction) int {
		return a.Date.Compare(b.Date)
	})

	balance := decimal.NewFromInt(0)
	for i, transaction := range transactions {
		switch transaction.Type {
		case TransactionTypeWithdraw:
			balance = balance.Sub(transaction.Amount)
		case TransactionTypeDeposit, TransactionTypeInterest:
			balance = balance.Add(transaction.Amount)
		default:
		}

		if balance.IsNegative() {
			return BankTransaction{}, errors.New("invalid balance")
		}

		transactions[i].Balance = balance.Copy()
	}

	// "commit" changes
	account := i.Accounts[params.AccountID]
	account.Transactions = transactions
	i.Accounts[params.AccountID] = account

	return newTransaction, nil
}

// NOTE: `date` has DAY level precision
func newBankTransactionID(transactions []BankTransaction, date time.Time) string {
	count := 1
	for _, transaction := range transactions {
		if transaction.Date.Year() == date.Year() &&
			transaction.Date.Month() == date.Month() &&
			transaction.Date.Day() == date.Day() {

			count++
		}
	}

	prefix := date.Format(DateFormatTransactionIDPrefix)
	suffix := fmt.Sprintf("%02d", count)

	return fmt.Sprintf("%s-%s", prefix, suffix)
}
