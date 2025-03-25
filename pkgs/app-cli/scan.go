package appcli

import (
	"bufio"
	"strings"
)

const PromptPrefix = "> "

func (a *AppCLI) Scan() (string, error) {
	a.Print(PromptPrefix)

	scanner := bufio.NewScanner(a.Reader)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	a.Println("")

	return strings.TrimSpace(scanner.Text()), scanner.Err()
}
