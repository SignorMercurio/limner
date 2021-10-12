package printer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SignorMercurio/limner/color"
)

// CustomPrinter prints data depending on custom func
type CustomPrinter struct {
	ColorPicker func(line string) color.Color
}

func (cp *CustomPrinter) Print(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintf(w, "%s\n", color.Apply(line, cp.ColorPicker(line)))
	}
}
