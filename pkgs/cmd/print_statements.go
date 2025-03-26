package cmd

import appctx "bank-acc-interest/pkgs/app-ctx"

type PrintStatements struct {
	*appctx.AppCtx
}

var _ Command = &PrintStatements{}

const MsgPrintStatementPrompt = "Please enter account and month to generate the statement <Account> <Year><Month>\n(or enter blank to go back to main menu):"

func (c *PrintStatements) Execute() {
	c.Println(MsgPrintStatementPrompt)

	// TODO: table format
	c.Print("| Period              | Num of days | EOD Balance | Rate Id | Rate | Annualized Interest      |")
}
