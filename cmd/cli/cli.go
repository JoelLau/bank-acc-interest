package main

import (
	appcli "bank-acc-interest/pkgs/app-cli"
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	logr := newLogger()

	app := appcli.NewAppCLI(os.Stdin, os.Stdout, logr)
	if err := app.Run(ctx); err != nil {
		logr.ErrorContext(ctx, "unexpected app error", slog.Any("error", err))
		return
	}
}

func newLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelDebug,
			},
		),
	)
}
