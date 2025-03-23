package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bank-acc-interest/pkgs/async"
	"bufio"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2EHelloNameFlow(t *testing.T) {
	t.Parallel()

	inputReader, inputWriter := io.Pipe()
	outputReader, outputWriter := io.Pipe()

	app := appcli.NewAppCLI(inputReader, outputWriter, discardLogger)
	ctx := t.Context()

	var wg sync.WaitGroup
	var appErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if appErr = app.Run(ctx); appErr != nil {
			require.NoError(t, appErr)
		}
	}()

	outputScanner := bufio.NewScanner(outputReader)

	var err error
	err = async.RunWithTimeOut(func() error {
		hasOutput := outputScanner.Scan()
		require.True(t, hasOutput)
		require.Equal(t, "what is your name?", outputScanner.Text())

		return nil
	}, 1*time.Second)
	require.NoError(t, err)

	_, err = inputWriter.Write([]byte("asdf\n"))
	require.NoError(t, err)

	err = async.RunWithTimeOut(func() error {
		hasOutput := outputScanner.Scan()
		require.True(t, hasOutput)
		require.Equal(t, "hello, asdf", outputScanner.Text())

		return nil
	}, 1*time.Second)
	require.NoError(t, err)

	wg.Wait()
}
