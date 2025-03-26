package appctx_test

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	t.Parallel()

	t.Run("single-line input", func(t *testing.T) {
		t.Parallel()

		var outputBuf bytes.Buffer

		want := "lorem ipsum"
		input := want + "\n"

		cli := appctx.NewAppCtx(strings.NewReader(input), &outputBuf, storage.NewInMemoryStorage())

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Equal(t, want, have)
	})

	t.Run("multi-line input", func(t *testing.T) {
		t.Parallel()

		var outputBuf bytes.Buffer

		input := "first\nsecond\n"
		cli := appctx.NewAppCtx(strings.NewReader(input), &outputBuf, storage.NewInMemoryStorage())

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Equal(t, "first", have)
	})

	t.Run("empty input", func(t *testing.T) {
		t.Parallel()

		var outputBuf bytes.Buffer

		cli := appctx.NewAppCtx(strings.NewReader(""), &outputBuf, storage.NewInMemoryStorage())

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Empty(t, have)
	})
}
