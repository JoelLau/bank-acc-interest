package appcli

import (
	"context"
	"fmt"
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
	menu := &Menu{a}

	err := menu.Run(ctx)
	if err != nil {
		err = fmt.Errorf("failed to run main menu: %w", err)
		return err
	}

	return nil
}
