package appcli_test

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// logs into the void
var discardLogger = slog.New(slog.NewTextHandler(io.Discard, nil))

var (
	ErrMockRead  = errors.New("mock read error")
	ErrMockWrite = errors.New("mock write error")
)

// mock io.Reader that always fails
type BrokenReader struct{}

// always returns error
func (f *BrokenReader) Read(p []byte) (n int, err error) {
	return 0, ErrMockRead
}

// mock io.Writer that always fails
type BrokenWriter struct{}

// always returns error
func (f *BrokenWriter) Write(p []byte) (n int, err error) {
	return n, ErrMockWrite
}

func TestBrokenWriter(t *testing.T) {
	t.Parallel()

	inputReader, _ := io.Pipe()

	app := appcli.NewAppCLI(inputReader, &BrokenWriter{}, discardLogger)
	ctx := t.Context()

	var wg sync.WaitGroup
	var appErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		appErr = app.Run(ctx)
	}()

	wg.Wait()
	require.ErrorIs(t, appErr, ErrMockWrite)
}

func TestBrokenReader(t *testing.T) {
	t.Parallel()

	var outputBuffer bytes.Buffer
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	app := appcli.NewAppCLI(&BrokenReader{}, &outputBuffer, logger)

	ctx := context.Background()

	err := app.Run(ctx)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrMockRead)
}
