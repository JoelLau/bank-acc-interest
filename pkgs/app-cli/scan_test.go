package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	t.Parallel()

	t.Run("single-line input", func(t *testing.T) {
		t.Parallel()

		var logBuffer bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&logBuffer, nil))

		want := "lorem ipsum"
		input := want + "\n"

		cli := appcli.NewAppCLI(strings.NewReader(input), nil, logger)

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Equal(t, want, have)
	})

	t.Run("multi-line input", func(t *testing.T) {
		t.Parallel()

		var logBuffer bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&logBuffer, nil))

		input := "first\nsecond\n"
		cli := appcli.NewAppCLI(strings.NewReader(input), nil, logger)

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Equal(t, "first", have)
	})

	t.Run("empty input", func(t *testing.T) {
		t.Parallel()

		var logBuffer bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&logBuffer, nil))

		cli := appcli.NewAppCLI(strings.NewReader(""), nil, logger)

		have, err := cli.Scan()
		require.NoError(t, err)
		require.Empty(t, have)
	})

	t.Run("reader error", func(t *testing.T) {
		t.Parallel()

		var logBuffer bytes.Buffer
		logger := slog.New(slog.NewTextHandler(&logBuffer, nil))

		cli := appcli.NewAppCLI(&BrokenReader{}, nil, logger)

		_, err := cli.Scan()
		require.ErrorIs(t, err, ErrMockRead)
	})
}
