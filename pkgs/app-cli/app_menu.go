package appcli

import (
	"context"
	"fmt"
)

type Menu struct {
	*AppCLI
}

const (
	MsgWelcomePrompt      = "Welcome to AwesomeGIC Bank! What would you like to do?"
	MsgAnythingElsePrompt = "\nIs there anything else you'd like to do?"
	MsgMenuItems          = "[T] Input transactions\n[I] Define interest rules\n[P] Print statement\n[Q] Quit"
	MsgExitThankyou       = "Thank you for banking with AwesomeGIC Bank.\nHave a nice day!"
)

func (a *Menu) Run(ctx context.Context) error {
	pretext := MsgWelcomePrompt

	for keepLooping := true; keepLooping; pretext = MsgAnythingElsePrompt {
		prompt := fmt.Sprintf("%s\n%s\n", pretext, MsgMenuItems)
		err := a.Println(prompt)
		if err != nil {
			err = fmt.Errorf("failed to print prompt: %w", err)
			return err
		}

		input, err := a.Scan()
		if err != nil {
			err = fmt.Errorf("failed to get user input for main menu: %w", err)
			return err
		}

		switch input {
		case "t", "T":
			inputTx := &InputTransactions{a.AppCLI}

			err = inputTx.Run(ctx)
			if err != nil {
				err = fmt.Errorf("failed to run input transaction: %w", err)
				return err
			}
		case "i", "I":
			// do nothing
		case "p", "P":
			// do nothing
		case "q", "Q":
			err = a.Println(MsgExitThankyou)
			if err != nil {
				err = fmt.Errorf("failed to print exit message: %w", err)
				return err
			}
			keepLooping = false
		default:
			keepLooping = false
		}
	}

	return nil
}
