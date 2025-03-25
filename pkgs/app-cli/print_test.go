package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appcli.NewAppCLI(inputReader, &outBuf, discardLogger)

	app.Print("asdf")

	stutter()
	require.Equal(t, outBuf.String(), "asdf")
	outBuf.Reset()
}

func TestPrintf(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appcli.NewAppCLI(inputReader, &outBuf, discardLogger)

	app.Printf("asdf %s", "qwer")

	stutter()
	require.Equal(t, outBuf.String(), "asdf qwer")
	outBuf.Reset()
}

func TestPrintln(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appcli.NewAppCLI(inputReader, &outBuf, discardLogger)

	app.Println("asdf")

	stutter()
	require.Equal(t, outBuf.String(), "asdf\n")
	outBuf.Reset()
}
