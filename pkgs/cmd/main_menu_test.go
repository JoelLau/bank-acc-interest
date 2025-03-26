package cmd_test

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/cmd"
	"bank-acc-interest/pkgs/storage"
	"bytes"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type SpyCommand struct {
	IsExecuted bool
}

var _ cmd.Command = &SpyCommand{}

func (c *SpyCommand) Execute() {
	c.IsExecuted = true
}

func TestE2EMainMenu(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, inputWriter := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	appCtx := appctx.NewAppCtx(inputReader, &outBuf, storage.NewInMemoryStorage())
	mainMenu := cmd.NewMainMenuCmd(appCtx)

	inputTxSpyCmd := &SpyCommand{}
	mainMenu.InputTransactions = inputTxSpyCmd
	require.Falsef(t, inputTxSpyCmd.IsExecuted, "sanity check: spy hasn't logged execution")

	defineInterestRuleSpyCmd := &SpyCommand{}
	mainMenu.DefineInterestRules = defineInterestRuleSpyCmd
	require.Falsef(t, defineInterestRuleSpyCmd.IsExecuted, "sanity check: spy hasn't logged execution")

	printTransactionsSpyCmd := &SpyCommand{}
	mainMenu.PrintStatements = printTransactionsSpyCmd
	require.Falsef(t, printTransactionsSpyCmd.IsExecuted, "sanity check: spy hasn't logged execution")

	var wg sync.WaitGroup
	var appErr, err error
	defer require.NoError(t, appErr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		mainMenu.Execute()
	}()

	welcomeMessage := `Welcome to AwesomeGIC Bank! What would you like to do?
[T] Input transactions
[I] Define interest rules
[P] Print statement
[Q] Quit`

	stutter()
	require.Contains(t, outBuf.String(), welcomeMessage)
	outBuf.Reset()

	stutter()
	_, err = inputWriter.Write([]byte("t\n"))
	require.NoError(t, err)
	require.True(t, inputTxSpyCmd.IsExecuted)

	stutter()
	_, err = inputWriter.Write([]byte("i\n"))
	require.NoError(t, err)
	require.True(t, defineInterestRuleSpyCmd.IsExecuted)

	stutter()
	_, err = inputWriter.Write([]byte("p\n"))
	require.NoError(t, err)
	require.True(t, printTransactionsSpyCmd.IsExecuted)

	stutter()
	_, err = inputWriter.Write([]byte("q\n"))
	require.NoError(t, err)

	thankyouMessage := `Thank you for banking with AwesomeGIC Bank.
Have a nice day!`

	stutter()
	require.Contains(t, outBuf.String(), thankyouMessage)
	outBuf.Reset()

	wg.Wait()
}

// short wait to allow app to write to buffer and simulate user input
func stutter() {
	time.Sleep(10 * time.Millisecond)
}
