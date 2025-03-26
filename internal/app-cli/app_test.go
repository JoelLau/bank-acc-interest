package appcli_test

import (
	appcli "bank-acc-interest/internal/app-cli"
	"bank-acc-interest/pkgs/storage"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppCLI(t *testing.T) {
	t.Parallel()

	var i io.Reader = &bytes.Buffer{}
	var o io.Writer = &bytes.Buffer{}
	var s storage.Storage = storage.NewInMemoryStorage()

	appCLI := appcli.NewAppCLI(i, o, s)
	appCLI.Run()

	require.Same(t, i, appCLI.AppCtx.Input)
	require.Same(t, o, appCLI.AppCtx.Output)
	require.Same(t, s, appCLI.AppCtx.Storage)
}
