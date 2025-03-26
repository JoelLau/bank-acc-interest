package cmd

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type InputTransactions struct {
	*appctx.AppCtx
}

var _ Command = &InputTransactions{}

const MsgInputTxPrompt = "Please enter transaction details in <Date> <Account> <Type> <Amount> format\n(or enter blank to go back to main menu):"

var TxColDef = []appctx.ColDef{
	{Header: "Date", Align: appctx.ColumnAlignLeft},
	{Header: "Txn Id", Align: appctx.ColumnAlignLeft},
	{Header: "Type", Align: appctx.ColumnAlignLeft},
	{Header: "Amount", Align: appctx.ColumnAlignRight},
}

func (c *InputTransactions) Execute() {

	for keepLooping := true; keepLooping; {
		c.Println(MsgInputTxPrompt)

		input, _ := c.Scan()

		switch input {
		case "":
			keepLooping = false
		default:
			tx, err := ParseInsertBankTxParams(input)
			if err != nil {
				c.Println("invalid input!\n")
				continue
			}

			_, err = c.Storage.InsertBankTransaction(tx)
			if err != nil {
				slog.Error(err.Error())
				c.Println("could not append bank transaction record\n")
				continue
			}

			transactions, err := c.Storage.GetAccountTransactions(tx.AccountID)
			if err != nil {
				c.Println("could not fetch transactions\n")
				continue
			}

			c.Printf("Account: %s\n", tx.AccountID)

			data := [][]string{}
			for _, transaction := range transactions {
				data = append(data, []string{
					transaction.Date.Format(DateFormatUserInput),
					transaction.ID,
					string(transaction.Type),
					transaction.Amount.StringFixed(2),
				})
			}

			c.PrintTable(TxColDef, data)
			c.Println("")

			keepLooping = false
		}
	}

	return
}

// Expects a string in "<Date> <Account> <Type> <Amount>" format
// Enhancement: refactor to break up parsers and validators - chain of responsbility pattern would be nice
func ParseInsertBankTxParams(s string) (storage.InsertBankTransactionParams, error) {
	var tx storage.InsertBankTransactionParams

	fields := strings.Fields(s)
	if len(fields) < 4 {
		return tx, ErrInvalidInput
	}

	date, err := time.Parse(DateFormatUserInput, fields[0])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction date: %w: %w", err, ErrInvalidInput)
		return storage.InsertBankTransactionParams{}, err
	}

	accID := strings.TrimSpace(fields[1])

	typStr := fields[2]
	var ttype storage.TransactionType

	switch typStr {
	case "w", "W":
		ttype = storage.TransactionTypeWithdraw
	case "d", "D":
		ttype = storage.TransactionTypeDeposit
	default:
		err = fmt.Errorf("invalid transaction type '%s': %w: %w", ttype, err, ErrInvalidInput)
		return storage.InsertBankTransactionParams{}, err
	}

	amt, err := decimal.NewFromString(fields[3])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction amount: %w: %w", err, ErrInvalidInput)
		return storage.InsertBankTransactionParams{}, err
	}
	if amt.IsNegative() {
		err := fmt.Errorf("%w: negative amount", ErrInvalidInput)
		return storage.InsertBankTransactionParams{}, err
	}
	if !amt.Round(2).Equal(amt) {
		err := fmt.Errorf("%w: too many decimal places", ErrInvalidInput)
		return storage.InsertBankTransactionParams{}, err
	}

	return storage.InsertBankTransactionParams{
		Date:      date,
		AccountID: accID,
		Type:      ttype,
		Amount:    amt.Round(2),
	}, nil
}
