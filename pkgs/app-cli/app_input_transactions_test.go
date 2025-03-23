package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestE2EInputTransactions(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, inputWriter := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appcli.NewAppCLI(inputReader, &outBuf, discardLogger)
	menu := appcli.InputTransactions{app}
	ctx := t.Context()

	var wg sync.WaitGroup
	var appErr, err error
	defer require.NoError(t, appErr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		appErr = menu.Run(ctx)
	}()

	inputTxPrompt := `Please enter transaction details in <Date> <Account> <Type> <Amount> format
(or enter blank to go back to main menu):
`

	stutter()
	require.Contains(t, outBuf.String(), inputTxPrompt)
	outBuf.Reset()

	stutter()
	_, err = inputWriter.Write([]byte("20230626 AC001 W 100.00\n"))
	require.NoError(t, err)

	stutter()
	require.Contains(t, outBuf.String(), "Account: AC001\n| Date     | Txn Id      | Type | Amount |\n| 20230626 | 20230626-02 | W    | 100.00 |\n")
	outBuf.Reset()

	wg.Wait()
}
