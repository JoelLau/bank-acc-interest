package appcli

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/cmd"
	"io"
)

// Wraps this app to facilitate tests.
type AppCLI struct {
	*appctx.AppCtx
}

func NewAppCLI(i io.Reader, o io.Writer) *AppCLI {
	return &AppCLI{
		AppCtx: appctx.NewAppCtx(i, o),
	}
}

func (a *AppCLI) Run() {
	menuCmd := cmd.MainMenu{AppCtx: a.AppCtx}
	menuCmd.Execute()
}
