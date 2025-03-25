package interestrule

import (
	"time"

	decimal "github.com/jackc/pgx-shopspring-decimal"
)

type InterestRuleID = string

// only one for each day should exists.
type InterestRule struct {
	// precise only up to "day" level.
	Date time.Time

	// string, free format.
	ID InterestRuleID

	// must be greater than 0 and less than 100.
	Rate decimal.Decimal
}
