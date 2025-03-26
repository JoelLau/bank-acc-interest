package storage

import (
	"errors"

	"github.com/shopspring/decimal"
)

type (
	UpsertInterestRuleParams = InterestRule
	DateString               = string
)

var ErrInvalidInput = errors.New("invalid input")

func (i *InMemoryStorage) UpsertInterestRule(params UpsertInterestRuleParams) (InterestRule, error) {
	if params.InterestRate.LessThan(decimal.NewFromInt(0)) {
		return InterestRule{}, ErrInvalidInput
	}

	if params.InterestRate.GreaterThan(decimal.NewFromInt(100)) {
		return InterestRule{}, ErrInvalidInput
	}

	rules := make([]InterestRule, len(i.InterestRules))
	copy(rules, i.InterestRules)

	newRule := InterestRule{
		Date:         params.Date,
		InterestRate: params.InterestRate,
		RuleID:       params.RuleID,
	}

	rules = append(rules, newRule)

	ruleMap := make(map[DateString]InterestRule)
	for _, rule := range rules {
		key := rule.Date.Format(DateFormatRuleMapKey)
		ruleMap[key] = rule
	}

	rules = make([]InterestRule, 0)
	for _, rule := range ruleMap {
		rules = append(rules, rule)
	}

	i.InterestRules = rules

	return newRule, nil
}
