package main

import (
	appcli "bank-acc-interest/internal/app-cli"
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"os"
)

func main() {
	app := appcli.AppCLI{
		AppCtx: appctx.NewAppCtx(
			os.Stdin,
			os.Stdout,
			storage.NewInMemoryStorage(),
		),
	}

	app.Run()
}
