package appcli_test

import (
	appcli "bank-acc-interest/internal/app-cli"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppCLI(t *testing.T) {
	t.Parallel()

	var i io.Reader = &bytes.Buffer{}
	var o io.Writer = &bytes.Buffer{}

	appCLI := appcli.NewAppCLI(i, o)
	appCLI.Run()

	require.Same(t, i, appCLI.AppCtx.Input)
	require.Same(t, o, appCLI.AppCtx.Output)
}
