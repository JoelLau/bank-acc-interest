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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrintStatementsExec(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, inputWriter := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	store := storage.NewInMemoryStorage()
	require.Len(t, store.InterestRules, 0)

	appCtx := appctx.NewAppCtx(inputReader, &outBuf, store)
	printStatementsCmd := cmd.PrintStatements{AppCtx: appCtx}

	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		printStatementsCmd.Execute()
	}()

	const msgPrompt = `Please enter account and month to generate the statement <Account> <Year><Month>
(or enter blank to go back to main menu):`

	stutter()
	require.Contains(t, outBuf.String(), msgPrompt)
	outBuf.Reset()

	stutter()
	_, err = inputWriter.Write([]byte("AC001 202306\n"))
	require.NoError(t, err)

	stutter()
	require.Contains(t, outBuf.String(), "Account: AC001\n| Period              | Num of days | EOD Balance | Rate Id | Rate | Annualized Interest      |")
	outBuf.Reset()
}

func TestParsePrintStatementParams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		given  string
		expect cmd.PrintStatementParams
		err    error
	}{
		{
			name:  "Example from REQUIREMENTS.md",
			given: "AC001 202306",
			expect: cmd.PrintStatementParams{
				AccountID: "AC001",
				Date:      time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			err: nil,
		},
		{
			name:   "Missing Tokens",
			given:  "AC001",
			expect: cmd.PrintStatementParams{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Invalid Date",
			given:  "AC001 202313",
			expect: cmd.PrintStatementParams{},
			err:    cmd.ErrInvalidInput,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			have, err := cmd.ParsePrintStatementParams(tt.given)
			require.ErrorIs(t, err, tt.err)

			assert.Equal(t, tt.expect.AccountID, have.AccountID)
			assert.Equal(t, tt.expect.Date, have.Date)
		})
	}
}
