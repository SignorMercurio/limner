package printer

import (
	"fmt"
	"io"

	"github.com/SignorMercurio/limner/color"
)

type SingleColorPrinter struct {
	Color color.Color
}

func (scp *SingleColorPrinter) Print(buf string, w io.Writer) {
	fmt.Fprintln(w, color.Apply(buf, scp.Color))
}
