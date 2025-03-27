package storage_test

import (
	"bank-acc-interest/pkgs/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInterestRules(t *testing.T) {
	t.Parallel()

	store := storage.NewInMemoryStorage()
	store.InterestRules = make([]storage.InterestRule, 5)
	require.Len(t, store.InterestRules, 5)

	rules, err := store.GetInterestRules()
	require.NoError(t, err)
	require.Len(t, rules, 5)
}
