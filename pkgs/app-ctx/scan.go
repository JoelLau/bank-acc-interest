package appctx

import (
	"bufio"
	"strings"
)

const PromptPrefix = "> "

func (a *AppCtx) Scan() (string, error) {
	a.Print(PromptPrefix)

	scanner := bufio.NewScanner(a.Input)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	a.Println("")

	return strings.TrimSpace(scanner.Text()), scanner.Err()
}
