package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	t.Parallel()

	t.Run("print", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		cli := appcli.NewAppCLI(nil, &buf, discardLogger)

		someText := "lorem ipsum"

		err := cli.Println(someText)
		require.NoError(t, err)
		require.Equal(t, someText+"\n", buf.String())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		cli := appcli.NewAppCLI(nil, &BrokenWriter{}, discardLogger)

		err := cli.Print("bad writer always fails")
		require.ErrorIs(t, err, ErrMockWrite)
	})
}
