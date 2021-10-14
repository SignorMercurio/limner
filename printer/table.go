package printer

import (
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

type TablePrinter struct {
	ColorPicker func(index int, column string) (color.Color, bool)
	firstLine   bool
}

func (tp *TablePrinter) Print(buf string, w io.Writer) {
	for _, line := range strings.Split(buf, "\n") {
		if tp.isHeader(line) {
			fmt.Fprintln(w, color.Apply(line, HeaderColor))
			tp.firstLine = false
			continue
		}

		tp.printTable(w, line, ColumnColors)
	}
}

func NewTablePrinter(colorPicker func(int, string) (color.Color, bool)) *TablePrinter {
	return &TablePrinter{
		ColorPicker: colorPicker,
		firstLine:   true,
	}
}

// isHeader returns if a line is a table header
func (tp *TablePrinter) isHeader(line string) bool {
	// --- --- also counts as headers
	allCapital := strings.ToUpper(line) == line

	return tp.firstLine || allCapital
}

// printTable prints a line to writer w as a row in the table
func (tp *TablePrinter) printTable(w io.Writer, line string, colors []color.Color) {
	columns := spaces.Split(line, -1)
	spaceIndices := spaces.FindAllStringIndex(line, -1)

	for i, column := range columns {
		c := colors[i&1]
		if tp.ColorPicker != nil {
			if cc, ok := tp.ColorPicker(i, column); ok {
				c = cc
			}
		}
		fmt.Fprint(w, color.Apply(column, c))

		// Write the spaces in between
		if i <= len(spaceIndices)-1 {
			spaceIndex := spaceIndices[i]
			fmt.Fprint(w, toSpaces(spaceIndex[1]-spaceIndex[0]))
		}
	}

	fmt.Fprintln(w)
}
