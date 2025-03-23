package async_test

import (
	"bank-acc-interest/pkgs/async"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRunWithTimeOut(t *testing.T) {
	t.Parallel()

	t.Run("fast functions complete", func(t *testing.T) {
		t.Parallel()

		fastFn := func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		}

		err := async.RunWithTimeOut(fastFn, 200*time.Millisecond)
		require.NoError(t, err)
	})

	t.Run("slow functions time out", func(t *testing.T) {
		t.Parallel()

		slowFn := func() error {
			time.Sleep(500 * time.Millisecond)
			return nil
		}

		err := async.RunWithTimeOut(slowFn, 100*time.Millisecond)
		require.ErrorIs(t, err, async.ErrTimeOut)
	})

	t.Run("function returns error", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")
		errorFn := func() error {
			return expectedErr
		}

		err := async.RunWithTimeOut(errorFn, 200*time.Millisecond)
		require.ErrorIs(t, err, expectedErr)
	})
}
