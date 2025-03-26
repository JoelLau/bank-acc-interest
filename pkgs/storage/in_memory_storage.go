package storage

import (
	"fmt"
	"time"
)

type InMemoryStrorage struct {
	// Enhancement: optimise - consider using a hash map + stack here
	BankTransactions []BankTransaction

	dateCounters map[DateString]int
}

var _ Storage = &InMemoryStrorage{}

type DateString = string

func NewInMemoryStorage() *InMemoryStrorage {
	return &InMemoryStrorage{
		BankTransactions: make([]BankTransaction, 0),
		dateCounters:     make(map[DateString]int),
	}
}

const DateFormatBankTx = "20060102"

func (i *InMemoryStrorage) InsertBankTransaction(params InsertBankTransactionParams) (BankTransaction, error) {
	// get "auto-increment" portion of id
	datestr := params.Date.Format(DateFormatBankTx)
	x, ok := i.dateCounters[datestr]
	if !ok {
		i.dateCounters[datestr] = 1
		x = i.dateCounters[datestr]
	}
	i.dateCounters[datestr] = i.dateCounters[datestr] + 1

	bankTx := BankTransaction{
		ID:        fmt.Sprintf("%s-%02d", datestr, x),
		CreatedAt: time.Now(),
		AccountID: params.AccountID,
		Date:      params.Date,
		Amount:    params.Amount,
		Type:      params.Type,
	}

	i.BankTransactions = append(i.BankTransactions, bankTx)
	return bankTx, nil
}
