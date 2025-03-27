package storage

func (i *InMemoryStrorage) GetInterestRules() ([]InterestRule, error) {
	return i.InterestRules, nil
}
