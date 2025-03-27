package storage

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (i *InMemoryStrorage) UpsertInterestRule(params UpsertInterestRuleParams) (InterestRule, error) {
	if params.InterestRate.LessThan(decimal.NewFromInt(0)) {
		return InterestRule{}, errors.New("invalid input")
	}

	if params.InterestRate.GreaterThan(decimal.NewFromInt(100)) {
		return InterestRule{}, errors.New("invalid input")
	}

	rules := make([]InterestRule, len(i.InterestRules))
	copy(rules, i.InterestRules)

	newRule := InterestRule{}
	rules = append(rules, newRule)

	ruleMap := make(map[RuleID]InterestRule)
	for _, rule := range rules {
		ruleMap[rule.RuleID] = rule
	}

	rules = make([]InterestRule, 0)
	for _, rule := range ruleMap {
		rules = append(rules, rule)
	}

	i.InterestRules = rules

	return newRule, nil
}

type UpsertInterestRuleParams struct {
	InterestRate decimal.Decimal
}

type RuleID = string
