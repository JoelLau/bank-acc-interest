package storage

type InMemoryStrorage struct {
	Accounts      map[AccountID]Account
	InterestRules []InterestRule
}

var _ Storage = &InMemoryStrorage{}

func NewInMemoryStorage() *InMemoryStrorage {
	return &InMemoryStrorage{
		Accounts:      make(map[AccountID]Account),
		InterestRules: make([]InterestRule, 0),
	}
}
