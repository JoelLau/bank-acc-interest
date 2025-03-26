package storage

func (i *InMemoryStorage) GetInterestRules() ([]InterestRule, error) {
	return i.InterestRules, nil
}
