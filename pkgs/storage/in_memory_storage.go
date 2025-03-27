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

// TODO: check account balance
func (i *InMemoryStrorage) InsertInterestRule(params InsertInterestRuleParams) (InterestRule, error) {
	rule := InterestRule(params)
	i.InterestRules = append(i.InterestRules, rule)
	return rule, nil
}
