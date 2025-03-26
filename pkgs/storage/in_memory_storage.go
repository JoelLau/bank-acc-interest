package storage

type InMemoryStorage struct {
	Accounts      map[AccountID]Account
	InterestRules []InterestRule
}

var _ Storage = &InMemoryStorage{}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		Accounts:      make(map[AccountID]Account),
		InterestRules: make([]InterestRule, 0),
	}
}
