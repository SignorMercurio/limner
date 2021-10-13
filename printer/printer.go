package printer

import (
	"io"
	"regexp"
)

var spaces = regexp.MustCompile(`\s{2,}`)

// Printer reads from r and writes to w
type Printer interface {
	Print(buf string, w io.Writer)
}
