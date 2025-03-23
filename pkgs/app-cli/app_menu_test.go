package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2EExit(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, inputWriter := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appcli.NewAppCLI(inputReader, &outBuf, discardLogger)
	menu := appcli.Menu{app}
	ctx := t.Context()

	var wg sync.WaitGroup
	var appErr, err error
	defer require.NoError(t, appErr)

	wg.Add(1)
	go func() {
		defer wg.Done()
		appErr = menu.Run(ctx)
	}()

	welcomeMessage := `Welcome to AwesomeGIC Bank! What would you like to do?
[T] Input transactions
[I] Define interest rules
[P] Print statement
[Q] Quit`

	stutter()
	require.Contains(t, outBuf.String(), welcomeMessage)
	outBuf.Reset()

	stutter()
	_, err = inputWriter.Write([]byte("q\n"))
	require.NoError(t, err)

	thankyouMessage := `Thank you for banking with AwesomeGIC Bank.
Have a nice day!`

	stutter()
	require.Contains(t, outBuf.String(), thankyouMessage)
	outBuf.Reset()

	wg.Wait()
}

// short wait to allow app to write to buffer and simulate user input
func stutter() {
	time.Sleep(100 * time.Millisecond)
}
