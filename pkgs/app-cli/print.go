package appcli

import (
	"fmt"
	"log/slog"
)

func (a *AppCLI) Print(s string) error {
	defer func() {
		p := recover()
		slog.Error(fmt.Sprintf("%s", p))
	}()

	_, err := a.Writer.Write([]byte(s))
	if err != nil {
		err = fmt.Errorf("failed to print '%s': %w", s, err)
		return err
	}

	return nil
}

func (a *AppCLI) Println(s string) error {
	return a.Print(s + "\n")
}
