package printer

import (
	"fmt"
	"io"
)

type PlainPrinter struct {
}

func (pp *PlainPrinter) Print(buf string, w io.Writer) {
	fmt.Fprint(w, buf)
}
