package cmd

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type DefineInterestRule struct {
	*appctx.AppCtx
}

var _ Command = &DefineInterestRule{}

const MsgDefineInterestRulePrompt = "Please enter interest rules details in <Date> <RuleId> <Rate in %> format\n(or enter blank to go back to main menu):"

func (c *DefineInterestRule) Execute() {

	for keepLooping := true; keepLooping; {
		c.Println(MsgDefineInterestRulePrompt)

		input, _ := c.Scan()

		switch input {
		case "":
			keepLooping = false
		default:
			tx, err := ParseInsertInterestRuleParams(input)
			if err != nil {
				c.Println("invalid input!\n")
				continue
			}

			_, err = c.Storage.InsertInterestRule(tx)
			if err != nil {
				c.Println("could not upsert interest rule record\n")
				continue
			}

			// TODO: print proper table
			c.Println("Interest rules:\n| Date     | RuleId | Rate (%) |")
			keepLooping = false
		}
	}

	return
}

// Expects a string in "<Date> <Account> <Type> <Amount>" format
// Enhancement: refactor to break up parsers and validators - chain of responsbility pattern would be nice
func ParseInsertInterestRuleParams(s string) (storage.InsertInterestRuleParams, error) {
	var tx storage.InsertInterestRuleParams

	fields := strings.Fields(s)
	if len(fields) < 3 {
		return tx, ErrInvalidInput
	}

	date, err := time.Parse(DateFormatUserInput, fields[0])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction date: %w: %w", err, ErrInvalidInput)
		return storage.InsertInterestRuleParams{}, err
	}

	ruleID := strings.TrimSpace(fields[1])

	interestRate, err := decimal.NewFromString(fields[2])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction amount: %w: %w", err, ErrInvalidInput)
		return storage.InsertInterestRuleParams{}, err
	}
	if interestRate.IsNegative() {
		err := fmt.Errorf("%w: negative amount", ErrInvalidInput)
		return storage.InsertInterestRuleParams{}, err
	}
	if !interestRate.Round(2).Equal(interestRate) {
		err := fmt.Errorf("%w: too many decimal places", ErrInvalidInput)
		return storage.InsertInterestRuleParams{}, err
	}

	return storage.InsertInterestRuleParams{
		Date:         date,
		RuleID:       ruleID,
		InterestRate: interestRate,
	}, nil
}
