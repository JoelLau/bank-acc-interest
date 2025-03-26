package cmd_test

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/cmd"
	"bank-acc-interest/pkgs/storage"
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintStatementsExec(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	store := storage.NewInMemoryStorage()
	require.Len(t, store.InterestRules, 0)

	appCtx := appctx.NewAppCtx(inputReader, &outBuf, store)
	printStatementsCmd := cmd.PrintStatements{AppCtx: appCtx}

	var wg sync.WaitGroup
	// var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		printStatementsCmd.Execute()
	}()

	const msgPrompt = `Please enter account and month to generate the statement <Account> <Year><Month>
(or enter blank to go back to main menu):`

	stutter()
	require.Contains(t, outBuf.String(), msgPrompt)
	require.Contains(t, outBuf.String(), "| Period              | Num of days | EOD Balance | Rate Id | Rate | Annualized Interest      |")
	outBuf.Reset()
}
