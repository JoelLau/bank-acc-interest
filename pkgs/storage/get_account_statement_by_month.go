package storage

import (
	"slices"
	"time"

	"github.com/shopspring/decimal"
)

// TODO: refactor
func (s *InMemoryStorage) GetAccountStatementByMonth(accountID AccountID, date time.Time) ([]BankTransaction, error) {
	transactions, _ := s.GetAccountTransactions(accountID)
	transactions = sortTransactionsByDateAsc(transactions)

	rules, _ := s.GetInterestRules()
	rules = sortRulesByDateAsc(rules)

	timePeriods := getTimePeriodsForMonth(transactions, rules, date)

	totalInterestEarned := decimal.NewFromInt(0)

	eodBalance := decimal.NewFromInt(0)

	for _, timePeriod := range timePeriods {
		eodBalance = decimal.NewFromInt(0)

		// TODO: optimise using doubly linked list
		for _, txn := range transactions {
			if !txn.Date.After(timePeriod.end) {
				switch txn.Type {
				case TransactionTypeWithdraw:
					eodBalance = eodBalance.Sub(txn.Amount)
				case TransactionTypeDeposit, TransactionTypeInterest:
					eodBalance = eodBalance.Add(txn.Amount)
				}
			}
		}

		// TODO: optimise using doubly linked list
		interestRatePerc := decimal.NewFromInt(0)
		for _, rule := range rules {
			if rule.Date.After(timePeriod.start) {
				break
			}
			interestRatePerc = rule.InterestRate.Copy()
		}

		daysInt := timePeriod.daysInRange()
		interestRate := interestRatePerc.Div(decimal.NewFromInt(100))

		interest := eodBalance.Mul(interestRate).Mul(decimal.NewFromInt(int64(daysInt)))

		totalInterestEarned = totalInterestEarned.Add(interest)
	}
	totalInterestEarned = totalInterestEarned.Div(decimal.NewFromInt(365))

	result := make([]BankTransaction, 0)
	for _, txn := range transactions {
		if txn.Date.Year() == date.Year() && txn.Date.Month() == date.Month() {
			result = append(result, txn)
		}
	}

	monthStart := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, -1)

	result = append(result, BankTransaction{
		Date:    monthEnd,
		ID:      "",
		Type:    TransactionTypeInterest,
		Amount:  totalInterestEarned,
		Balance: eodBalance.Add(totalInterestEarned),
	})

	return result, nil
}

func getTimePeriodsForMonth(transactions []BankTransaction, rules []InterestRule, date time.Time) []period {
	uniqueDates := make(map[time.Time]bool)

	monthStart := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, -1)

	uniqueDates[monthStart] = true

	for _, t := range transactions {
		d := time.Date(t.Date.Year(), t.Date.Month(), t.Date.Day(), 0, 0, 0, 0, time.UTC)
		uniqueDates[d] = true
	}

	for _, r := range rules {
		d := time.Date(r.Date.Year(), r.Date.Month(), r.Date.Day(), 0, 0, 0, 0, time.UTC)
		uniqueDates[d] = true
	}

	dateAsc := make([]time.Time, 0, len(uniqueDates))
	for d := range uniqueDates {
		if d.Year() == date.Year() && d.Month() == date.Month() {
			dateAsc = append(dateAsc, d)
		}
	}
	slices.SortFunc(dateAsc, func(a, b time.Time) int {
		return a.Compare(b)
	})

	periods := make([]period, 0)

	curr := dateAsc[0]
	for _, d := range dateAsc[1:] {
		periods = append(periods, period{
			start: curr,
			end:   d.AddDate(0, 0, -1),
		})
		curr = d
	}

	if curr.Before(monthEnd) || curr.Equal(monthEnd) {
		periods = append(periods, period{
			start: curr,
			end:   monthEnd,
		})
	}

	return periods
}

func sortTransactionsByDateAsc(transactions []BankTransaction) []BankTransaction {
	sorted := make([]BankTransaction, len(transactions))
	copy(sorted, transactions)

	slices.SortFunc(sorted, func(a, b BankTransaction) int {
		return a.Date.Compare(b.Date)
	})

	return sorted
}

func sortRulesByDateAsc(transactions []InterestRule) []InterestRule {
	sorted := make([]InterestRule, len(transactions))
	copy(sorted, transactions)

	slices.SortFunc(sorted, func(a, b InterestRule) int {
		return a.Date.Compare(b.Date)
	})

	return sorted
}

type period struct {
	start   time.Time
	end     time.Time
	balance decimal.Decimal // EOD balance
	rate    decimal.Decimal
}

func (p *period) daysInRange() int {
	return int(p.end.Sub(p.start).Hours()/24) + 1
}
