package main

import (
	appcli "bank-acc-interest/internal/app-cli"
	appctx "bank-acc-interest/pkgs/app-ctx"
	"os"
)

func main() {
	app := appcli.AppCLI{
		AppCtx: appctx.NewAppCtx(
			os.Stdin,
			os.Stdout,
		),
	}

	app.Run()
}
