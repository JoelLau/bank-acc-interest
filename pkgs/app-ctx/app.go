package appctx

import (
	"io"
)

// Utility class for simplifying app I/O
type AppCtx struct {
	Input  io.Reader
	Output io.Writer
}

func NewAppCtx(i io.Reader, o io.Writer) *AppCtx {
	return &AppCtx{
		Input:  i,
		Output: o,
	}
}
