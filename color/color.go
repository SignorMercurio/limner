package color

import "fmt"

type Color int

const (
	Black Color = 30 + iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	escapeStart = "\x1b["
	escapeEnd   = "\x1b[0m"
)

// Apply colorize line with color c
func Apply(line string, c Color) string {
	return fmt.Sprintf("%s%dm%s%s", escapeStart, int(c), line, escapeEnd)
}
