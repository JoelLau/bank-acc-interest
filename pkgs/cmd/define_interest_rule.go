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

var InterestRuleColDef = []appctx.ColDef{
	{Header: "Date", Align: appctx.ColumnAlignLeft},
	{Header: "RuleId", Align: appctx.ColumnAlignLeft},
	{Header: "Rate (%)", Align: appctx.ColumnAlignRight},
}

func (c *DefineInterestRule) Execute() {

	for keepLooping := true; keepLooping; {
		c.Println(MsgDefineInterestRulePrompt)

		input, _ := c.Scan()

		switch input {
		case "":
			keepLooping = false
		default:
			tx, err := ParseUpsertInterestRuleParams(input)
			if err != nil {
				c.Println("invalid input!\n")
				continue
			}

			_, err = c.Storage.UpsertInterestRule(tx)
			if err != nil {
				c.Println("could not upsert interest rule record\n")
				continue
			}

			rules, err := c.Storage.GetInterestRules()
			if err != nil {
				c.Println("could not get interest rules\n")
				continue
			}

			data := make([][]string, len(rules))
			for i, rule := range rules {
				data[i] = []string{
					rule.Date.Format(DateFormatUserInput),
					rule.RuleID,
					rule.InterestRate.StringFixed(2),
				}
			}

			c.PrintTable(InterestRuleColDef, data)
			keepLooping = false
		}
	}

	return
}

// Expects a string in "<Date> <Account> <Type> <Amount>" format
// Enhancement: refactor to break up parsers and validators - chain of responsbility pattern would be nice
func ParseUpsertInterestRuleParams(s string) (storage.UpsertInterestRuleParams, error) {
	var tx storage.UpsertInterestRuleParams

	fields := strings.Fields(s)
	if len(fields) < 3 {
		return tx, ErrInvalidInput
	}

	date, err := time.Parse(DateFormatUserInput, fields[0])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction date: %w: %w", err, ErrInvalidInput)
		return storage.UpsertInterestRuleParams{}, err
	}

	ruleID := strings.TrimSpace(fields[1])

	interestRate, err := decimal.NewFromString(fields[2])
	if err != nil {
		err = fmt.Errorf("failed to parse transaction amount: %w: %w", err, ErrInvalidInput)
		return storage.UpsertInterestRuleParams{}, err
	}
	if interestRate.IsNegative() {
		err := fmt.Errorf("%w: negative amount", ErrInvalidInput)
		return storage.UpsertInterestRuleParams{}, err
	}
	if !interestRate.Round(2).Equal(interestRate) {
		err := fmt.Errorf("%w: too many decimal places", ErrInvalidInput)
		return storage.UpsertInterestRuleParams{}, err
	}

	return storage.UpsertInterestRuleParams{
		Date:         date,
		RuleID:       ruleID,
		InterestRate: interestRate,
	}, nil
}
