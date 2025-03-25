package appctx

import (
	"fmt"
)

func (a *AppCtx) Printf(s string, args ...any) {
	fmt.Fprintf(a.Output, s, args...)
}

func (a *AppCtx) Print(s string) {
	fmt.Fprint(a.Output, s)
}

func (a *AppCtx) Println(s string) {
	fmt.Fprintln(a.Output, s)
}
