package appcli

import (
	"bufio"
	"strings"
)

func (a *AppCLI) Scan() (string, error) {
	scanner := bufio.NewScanner(a.Reader)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return strings.TrimSpace(scanner.Text()), scanner.Err()
}
