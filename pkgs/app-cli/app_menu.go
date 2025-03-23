package appcli

import (
	"context"
	"fmt"
)

type Menu struct {
	*AppCLI
}

const PromptMenu = `
Welcome to AwesomeGIC Bank! What would you like to do?
[T] Input transactions
[I] Define interest rules
[P] Print statement
[Q] Quit`

const ExitMessage = `
Thank you for banking with AwesomeGIC Bank.
Have a nice day!
`

func (a *Menu) Run(ctx context.Context) error {

	for keepLooping := true; keepLooping; {
		err := a.Print(PromptMenu)
		if err != nil {
			err = fmt.Errorf("failed to prompt main menu: %w", err)
			return err
		}

		input, err := a.Scan()
		if err != nil {
			err = fmt.Errorf("failed to get user input for main menu: %w", err)
			return err
		}

		switch input {
		case "q", "Q":
			err = a.Println(ExitMessage)
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
