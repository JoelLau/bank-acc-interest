package appctx

import (
	"fmt"
)

func (a *AppCtx) Printf(s string, args ...any) {
	_, _ = a.Output.Write([]byte(fmt.Sprintf(s, args...)))
}

func (a *AppCtx) Print(s string) {
	_, _ = a.Output.Write([]byte(s))
}

func (a *AppCtx) Println(s string) {
	_, _ = a.Output.Write([]byte(s + "\n"))
}
