package appcli

import (
	"fmt"
)

func (a *AppCLI) Print(s string) error {
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
