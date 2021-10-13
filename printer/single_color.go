package printer

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SignorMercurio/limner/color"
)

type SingleColorPrinter struct {
	Color color.Color
}

func (scp *SingleColorPrinter) Print(r io.Reader, w io.Writer) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		fmt.Fprintln(w, color.Apply(line, scp.Color))
	}
}
