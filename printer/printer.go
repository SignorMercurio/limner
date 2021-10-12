package printer

import (
	"io"
)

// Printer reads from r and writes to w
type Printer interface {
	Print(r io.Reader, w io.Writer)
}
