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

var PrintStmtColDef = []appctx.ColDef{
	{Header: "Period", Align: appctx.ColumnAlignLeft},
	{Header: "Txn Id", Align: appctx.ColumnAlignLeft},
	{Header: "Type", Align: appctx.ColumnAlignLeft},
	{Header: "Amount", Align: appctx.ColumnAlignRight},
	{Header: "Balance", Align: appctx.ColumnAlignRight},
}

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

			statements, err := c.Storage.GetAccountStatementByMonth(params.AccountID, params.Date)
			if err != nil {
				c.Println("error getting account statements\n")
				continue
			}

			data := make([][]string, len(statements))
			for i, stmt := range statements {
				data[i] = []string{
					stmt.Date.Format(DateFormatUserInput),
					stmt.ID,
					string(stmt.Type),
					stmt.Amount.StringFixed(2),
					stmt.Balance.StringFixed(2),
				}
			}

			c.Printf("Account: %s\n", params.AccountID)
			c.PrintTable(PrintStmtColDef, data)
			c.Println("")

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
