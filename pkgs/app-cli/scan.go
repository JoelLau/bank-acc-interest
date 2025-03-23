package appcli

import (
	"bufio"
	"strings"
)

const PromptPrefix = "> "

func (a *AppCLI) Scan() (string, error) {
	// if err := a.Print(PromptPrefix); err != nil {
	// 	return "", err
	// }

	scanner := bufio.NewScanner(a.Reader)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return strings.TrimSpace(scanner.Text()), scanner.Err()
}
