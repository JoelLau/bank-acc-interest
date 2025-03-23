package appcli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type AppCLI struct {
	Reader io.Reader
	Writer io.Writer
	Logger *slog.Logger
}

func NewAppCLI(r io.Reader, w io.Writer, l *slog.Logger) *AppCLI {
	return &AppCLI{Reader: r, Writer: w, Logger: l}
}

func (a *AppCLI) Run(ctx context.Context) error {
	_, err := a.Writer.Write([]byte("what is your name?\n"))
	if err != nil {
		err = fmt.Errorf("failed to prompt for name: %w", err)
		return err
	}

	scanner := bufio.NewScanner(a.Reader)
	scanner.Scan()

	name := strings.TrimSpace(scanner.Text())

	_, err = a.Writer.Write(fmt.Appendf([]byte{}, "hello, %s\n", name))
	if err != nil {
		err = fmt.Errorf("failed to print 'hello, <name>': %w", err)
		return err
	}

	return nil
}
