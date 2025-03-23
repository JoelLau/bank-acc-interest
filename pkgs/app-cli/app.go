package appcli

import (
	"context"
	"io"
	"log/slog"
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
	if err := a.Println("what is your name?"); err != nil {
		return err
	}

	name, err := a.Scan()
	if err != nil {
		return err
	}

	if err := a.Println("hello, " + name); err != nil {
		return err
	}

	return nil
}
