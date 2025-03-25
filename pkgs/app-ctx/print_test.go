package appctx_test

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf)

	app.Print("asdf")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf")
	outBuf.Reset()
}

func TestPrintf(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf)

	app.Printf("asdf %s", "qwer")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf qwer")
	outBuf.Reset()
}

func TestPrintln(t *testing.T) {
	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf)

	app.Println("asdf")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf\n")
	outBuf.Reset()
}
