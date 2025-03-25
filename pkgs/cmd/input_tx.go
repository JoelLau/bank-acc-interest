package cmd

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/transactions"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type InputTransactions struct {
	*appctx.AppCtx
}

var _ Command = &InputTransactions{}

const MsgInputTxPrompt = "Please enter transaction details in <Date> <Account> <Type> <Amount> format\n(or enter blank to go back to main menu):"

func (c *InputTransactions) Execute() {

	for keepLooping := true; keepLooping; {
		c.Println(MsgInputTxPrompt)

		input, _ := c.Scan()

		switch input {
		case "":
			keepLooping = false
		default:
			_, err := ParseTransactionString(input)
			if err != nil {
				c.Println("invalid input!\n")
				continue
			}

			// TODO: "persist" transaction
			// _, err = a.Repo.CreateTransaction(tx)
			// if err != nil {
			// 	a.Println("invalid input!\n")
			// 	continue
			// }

			// TODO: print table
			c.Println("Account: AC001\n| Date     | Txn Id      | Type | Amount |\n| 20230626 | 20230626-02 | W    | 100.00 |")
			keepLooping = false
		}
	}

	return
}

// TODO: move this
const DateFormatUserInput = "20060102"

// TODO: move this
var ErrInvalidInput = errors.New("invalid input")

// Expects a string in "<Date> <Account> <Type> <Amount>" format
// TODO: refactor to break up parsers and validators - chain of responsbility pattern would be nice
func ParseTransactionString(s string) (transactions.Transaction, error) {
	var tx transactions.Transaction

	fields := strings.Fields(s)
	if len(fields) < 4 {
		return tx, ErrInvalidInput
	}

	date, err := time.Parse(DateFormatUserInput, fields[0])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction date: %w: %w", err, ErrInvalidInput)
		return transactions.Transaction{}, err
	}

	accID := strings.TrimSpace(fields[1])

	ttype := fields[2]

	switch ttype {
	case "w", "W":
		ttype = "W"
	case "d", "D":
		ttype = "D"
	default:
		err = fmt.Errorf("invalid transaction type '%s': %w: %w", ttype, err, ErrInvalidInput)
		return transactions.Transaction{}, err
	}

	amt, err := decimal.NewFromString(fields[3])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction amount: %w: %w", err, ErrInvalidInput)
		return transactions.Transaction{}, err
	}
	if amt.IsNegative() {
		err := fmt.Errorf("%w: negative amount", ErrInvalidInput)
		return transactions.Transaction{}, err
	}
	if !amt.Round(2).Equal(amt) {
		err := fmt.Errorf("%w: too many decimal places", ErrInvalidInput)
		return transactions.Transaction{}, err
	}

	return transactions.Transaction{
		Date:      date,
		AccountID: accID,
		Type:      transactions.TransactionType(ttype),
		Amount:    amt.Round(2),
	}, nil
}
