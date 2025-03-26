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

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2EIDefineInterestRule(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, inputWriter := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	store := storage.NewInMemoryStorage()
	require.Len(t, store.InterestRules, 0)

	appCtx := appctx.NewAppCtx(inputReader, &outBuf, store)
	defineInterestRuleCmd := cmd.DefineInterestRule{AppCtx: appCtx}

	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		defineInterestRuleCmd.Execute()
	}()

	enterDetailsPrompt := `Please enter interest rules details in <Date> <RuleId> <Rate in %> format
(or enter blank to go back to main menu):`

	stutter()
	require.Contains(t, outBuf.String(), enterDetailsPrompt)
	outBuf.Reset()

	_, err = inputWriter.Write([]byte("gibberish\n"))
	require.NoError(t, err)

	stutter()
	require.Contains(t, outBuf.String(), "invalid input")
	outBuf.Reset()
	require.Len(t, store.InterestRules, 0)

	_, err = inputWriter.Write([]byte("20230615 RULE03 2.20\n"))
	require.NoError(t, err)

	stutter()
	require.Contains(t, outBuf.String(), `| Date     | RuleId | Rate (%) |
| 20230615 | RULE03 |     2.20 |`)
	outBuf.Reset()
	require.Len(t, store.InterestRules, 1)

	wg.Wait()
}

func TestParseUpsertInterestRuleParams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		given  string
		expect storage.InterestRule
		err    error
	}{
		{
			name:  "Example from REQUIREMENTS.md",
			given: "20230615 RULE03 2.20",
			expect: storage.InterestRule{
				Date:         time.Date(2023, time.June, 15, 0, 0, 0, 0, time.UTC),
				RuleID:       "RULE03",
				InterestRate: decimal.NewFromFloat(2.2),
			},
			err: nil,
		},
		{
			name:   "Invalid Date format",
			given:  "20230699 RULE03 2.20",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Amount must be > 0",
			given:  "20230699 RULE03 -2.20",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Amount must be < 100",
			given:  "20230699 RULE03 200.20",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "2 DP precision",
			given:  "20230615 RULE03 2.202",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Invalid Rate",
			given:  "20230615 RULE03 asdf",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Rate must be > 0",
			given:  "20230615 RULE03 -0.01",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Missing Fields",
			given:  "20230615 2.20",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
		{
			name:   "Extra Fields",
			given:  "20230615 RULE03 2.202 John Doe",
			expect: storage.InterestRule{},
			err:    cmd.ErrInvalidInput,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			have, err := cmd.ParseUpsertInterestRuleParams(tt.given)
			require.ErrorIs(t, err, tt.err)

			assert.Equal(t, tt.expect.Date, have.Date)
			assert.Equal(t, tt.expect.RuleID, have.RuleID)
			assert.True(t, tt.expect.InterestRate.Equal(have.InterestRate))
		})
	}
}
