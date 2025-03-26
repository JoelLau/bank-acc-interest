package appcli

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/cmd"
	"bank-acc-interest/pkgs/storage"
	"io"
)

// Wraps this app to facilitate tests.
type AppCLI struct {
	*appctx.AppCtx
}

func NewAppCLI(i io.Reader, o io.Writer, s storage.Storage) *AppCLI {
	return &AppCLI{
		AppCtx: appctx.NewAppCtx(i, o, s),
	}
}

func (a *AppCLI) Run() {
	menuCmd := cmd.NewMainMenuCmd(a.AppCtx)
	menuCmd.Execute()
}
