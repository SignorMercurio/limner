package printer

import (
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

// CustomPrinter prints data depending on custom func
type CustomPrinter struct {
	ColorPicker func(line string) color.Color
}

func (cp *CustomPrinter) Print(buf string, w io.Writer) {
	for _, line := range strings.Split(buf, "\n") {
		fmt.Fprintf(w, "%s\n", color.Apply(line, cp.ColorPicker(line)))
	}
}
