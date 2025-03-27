package cmd

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"fmt"
	"strings"
	"time"
)

type PrintStatements struct {
	*appctx.AppCtx
}

var _ Command = &PrintStatements{}

const MsgPrintStatementPrompt = "Please enter account and month to generate the statement <Account> <Year><Month>\n(or enter blank to go back to main menu):"
const DateFormatPrintStatement = "200601"

func (c *PrintStatements) Execute() {
	for keepLooping := true; keepLooping; {
		c.Println(MsgPrintStatementPrompt)

		input, _ := c.Scan()

		switch input {
		case "":
			keepLooping = false
		default:
			params, err := ParsePrintStatementParams(input)
			if err != nil {
				c.Println("invalid input")
				continue
			}

			// TODO: table format
			c.Printf("Account: %s\n", params.AccountID)
			c.Print("| Period              | Num of days | EOD Balance | Rate Id | Rate | Annualized Interest      |")
			keepLooping = false
		}
	}
}

type PrintStatementParams struct {
	AccountID storage.AccountID
	Date      time.Time // month level precision
}

func ParsePrintStatementParams(s string) (PrintStatementParams, error) {
	fields := strings.Fields(s)
	if len(fields) < 2 {
		return PrintStatementParams{}, ErrInvalidInput
	}

	accID := strings.TrimSpace(fields[0])

	date, err := time.Parse(DateFormatPrintStatement, fields[1])
	if err != nil {
		err = fmt.Errorf("failed to parse statement date: %w: %w", err, ErrInvalidInput)
		return PrintStatementParams{}, err
	}

	return PrintStatementParams{
		Date:      date,
		AccountID: accID,
	}, nil
}
