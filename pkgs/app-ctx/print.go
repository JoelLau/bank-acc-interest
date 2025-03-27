package appctx

import (
	"fmt"
	"strings"
)

func (a *AppCtx) Printf(s string, args ...any) {
	_, _ = a.Output.Write(fmt.Appendf(nil, s, args...))
}

func (a *AppCtx) Print(s string) {
	_, _ = a.Output.Write([]byte(s))
}

func (a *AppCtx) Println(s string) {
	_, _ = a.Output.Write([]byte(s + "\n"))
}

type ColDef struct {
	Header string // e.g. "Date", "Txn Id", etc
	Align  ColumnAlign
}

type ColumnAlign string

const (
	ColumnAlignLeft  ColumnAlign = "left"
	ColumnAlignRight ColumnAlign = "right"
)

func (a *AppCtx) PrintTable(colDef []ColDef, data [][]string) {
	colCount := len(colDef)

	// add coldef to data
	headers := make([]string, len(colDef))
	for idx, def := range colDef {
		headers[idx] = def.Header
	}
	data = append([][]string{headers}, data...)

	// get widths
	widths := make([]int, colCount)
	for _, row := range data {
		for x, cell := range row {
			widths[x] = max(widths[x], len(cell))
		}
	}

	sb := strings.Builder{}

	for y, row := range data {
		sb.WriteString("| ")
		for x, col := range row {
			if colDef[x].Align == ColumnAlignRight {
				sb.WriteString(fmt.Sprintf("%*s |", widths[x], col))
			} else {
				sb.WriteString(fmt.Sprintf("%-*s |", widths[x], col))
			}

			if x < len(row)-1 {
				sb.WriteString(" ")
			}
		}

		if y < len(data)-1 {
			sb.WriteString("\n")
		}
	}

	fmt.Fprint(a.Output, sb.String())
}
