package appcli

import (
	"fmt"
)

func (a *AppCLI) Printf(s string, args ...any) {
	fmt.Fprintf(a.Writer, s, args...)
}

func (a *AppCLI) Print(s string) {
	fmt.Fprint(a.Writer, s)
}

func (a *AppCLI) Println(s string) {
	fmt.Fprintln(a.Writer, s)
}
