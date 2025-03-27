package appctx_test

import (
	appctx "bank-acc-interest/pkgs/app-ctx"
	"bank-acc-interest/pkgs/storage"
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrint(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf, storage.NewInMemoryStorage())

	app.Print("asdf")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf")
	outBuf.Reset()
}

func TestPrintf(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf, storage.NewInMemoryStorage())

	app.Printf("asdf %s", "qwer")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf qwer")
	outBuf.Reset()
}

func TestPrintln(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf, storage.NewInMemoryStorage())

	app.Println("asdf")

	time.Sleep(10 * time.Millisecond)
	require.Equal(t, outBuf.String(), "asdf\n")
	outBuf.Reset()
}

func TestPrintTable(t *testing.T) {
	t.Parallel()

	// NOTE: sleep for short duration so that app.Run() can write to buffer
	inputReader, _ := io.Pipe()

	// NOTE: remember to run .Reset() after reading
	var outBuf bytes.Buffer

	app := appctx.NewAppCtx(inputReader, &outBuf, storage.NewInMemoryStorage())

	time.Sleep(10 * time.Millisecond)

	app.PrintTable([]appctx.ColDef{
		{Header: "Date", Align: appctx.ColumnAlignLeft},
		{Header: "Txn Id", Align: appctx.ColumnAlignLeft},
		{Header: "Type", Align: appctx.ColumnAlignLeft},
		{Header: "Amount", Align: appctx.ColumnAlignRight},
	}, [][]string{{"20230505", "20230505-01", "D", "20.00"}})

	expect := `| Date     | Txn Id      | Type | Amount |
| 20230505 | 20230505-01 | D    |  20.00 |`

	require.Equal(t, expect, outBuf.String())
	outBuf.Reset()
}
