package appctx

import (
	"bank-acc-interest/pkgs/storage"
	"io"
)

// Utility class for simplifying app I/O
type AppCtx struct {
	Input   io.Reader
	Output  io.Writer
	Storage storage.Storage
}

func NewAppCtx(i io.Reader, o io.Writer, s storage.Storage) *AppCtx {
	return &AppCtx{
		Input:   i,
		Output:  o,
		Storage: s,
	}
}
