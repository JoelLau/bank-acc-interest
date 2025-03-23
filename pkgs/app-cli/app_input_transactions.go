package appcli

import (
	"context"
	"fmt"
)

type InputTransactions struct {
	*AppCLI
}

const MsgInputTxPrompt = "Please enter transaction details in <Date> <Account> <Type> <Amount> format\n(or enter blank to go back to main menu):"

func (a *InputTransactions) Run(ctx context.Context) error {

	for keepLooping := true; keepLooping; {
		err := a.Println(MsgInputTxPrompt)
		if err != nil {
			err = fmt.Errorf("failed to prompt input transactions: %w", err)
			return err
		}

		input, err := a.Scan()
		if err != nil {
			err = fmt.Errorf("failed to get user input for input transactions: %w", err)
			return err
		}

		switch input {
		case "":
			keepLooping = false
		default:
			// TODO: parse, handle input
			// TODO: print the rest of the table
			err := a.Println(fmt.Sprintf("Account: %s\n| Date     | Txn Id      | Type | Amount |\n| 20230626 | 20230626-02 | W    | 100.00 |", "AC001"))
			if err != nil {
				err = fmt.Errorf("failed to input transaction response: %w", err)
				return err
			}
			keepLooping = false
		}
	}

	return nil
}
