package cmd

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"fmt"
)

type MainMenu struct {
	*appctx.AppCtx

	InputTransactions *InputTransactions
}

func NewMainMenuCmd(AppCtx *appctx.AppCtx) *MainMenu {
	return &MainMenu{
		AppCtx:            AppCtx,
		InputTransactions: &InputTransactions{AppCtx: AppCtx},
	}
}

var _ Command = &MainMenu{}

const (
	MsgWelcomePrompt      = "Welcome to AwesomeGIC Bank! What would you like to do?"
	MsgAnythingElsePrompt = "\nIs there anything else you'd like to do?"
	MsgMenuItems          = "[T] Input transactions\n[I] Define interest rules\n[P] Print statement\n[Q] Quit"
	MsgExitThankyou       = "Thank you for banking with AwesomeGIC Bank.\nHave a nice day!"
)

func (c *MainMenu) Execute() {
	pretext := MsgWelcomePrompt

	for keepLooping := true; keepLooping; pretext = MsgAnythingElsePrompt {
		c.Printf("%s\n%s\n", pretext, MsgMenuItems)

		input, err := c.Scan()
		if err != nil {
			err = fmt.Errorf("failed to get user input for main menu: %w", err)
			panic(err)
		}

		switch input {
		case "t", "T":
			c.InputTransactions.Execute()
		case "i", "I":
			// do nothing
		case "p", "P":
			// do nothing
		case "q", "Q":
			c.Println(MsgExitThankyou)
			keepLooping = false
		default:
			keepLooping = false
		}
	}

	return
}
