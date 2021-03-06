package printer

import (
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

type YamlPrinter struct {
	inString bool
	inTable  bool
}

func (yp *YamlPrinter) Print(buf string, w io.Writer) {
	for _, line := range strings.Split(buf, "\n") {
		yp.printYaml(line, w)
	}
}

func NewYamlPrinter() *YamlPrinter {
	return &YamlPrinter{}
}

// printYaml prints a line to writer w in yaml format
func (yp *YamlPrinter) printYaml(line string, w io.Writer) {
	indentCnt := getIndent(line)
	indent := toSpaces(indentCnt)
	line = strings.TrimLeft(line, " ")

	if yp.inString {
		// the line is a part of a string
		fmt.Fprintf(w, "%s%s\n", indent, yp.colorString(line))
		yp.inString = !yp.isClosed(line)
		return
	}

	// value could contain ": "
	splitted := strings.SplitN(line, ": ", 2)
	key := splitted[0]

	// normal key: value
	if len(splitted) == 2 {
		value := splitted[1]
		trimmedValue := strings.TrimLeft(value, " ")
		valueIndent := toSpaces(len(value) - len(trimmedValue))
		fmt.Fprintf(w, "%s%s: %s%s\n", indent, yp.colorKey(key), valueIndent, yp.colorValue(trimmedValue))
		yp.inString = yp.isOpen(trimmedValue)

		yp.inTable = false
		return
	}

	// key: OR - arrayItem
	if strings.HasSuffix(key, ":") {
		fmt.Fprintf(w, "%s%s\n", indent, yp.colorKey(key))

		yp.inTable = false
		return
	}

	// special case: table
	columns := spaces.Split(line, -1)
	if len(columns) >= 2 {
		if !yp.inTable {
			fmt.Fprintln(w, color.Apply(key, headerColor))
			yp.inTable = true
		} else {
			NewTablePrinter(ColorStatus).printTable(w, key, columnColors)
		}
		return
	}

	fmt.Fprintf(w, "%s%s\n", indent, yp.colorValue(key))
}

// isClosed returns if the string is closed in this line
func (yp *YamlPrinter) isClosed(line string) bool {
	return strings.HasSuffix(line, "'") || strings.HasSuffix(line, `"`)
}

// isOpen returns if the string is still open in this line
func (yp *YamlPrinter) isOpen(line string) bool {
	return (strings.HasPrefix(line, "'") && !strings.HasSuffix(line, "'")) || (strings.HasPrefix(line, `"`) && !strings.HasSuffix(line, `"`))
}

// isString returns if the line is double quoted
func (yp *YamlPrinter) isString(line string) bool {
	return strings.HasPrefix(line, `"`) && strings.HasSuffix(line, `"`)
}

// inArray returns if the line starts with "- "
func (yp *YamlPrinter) inArray(line string) bool {
	return strings.HasPrefix(line, "- ")
}

// colorString colorizes yaml strings
func (yp *YamlPrinter) colorString(line string) string {
	format := `%s`
	if yp.isString(line) {
		format = `"%s"`
	}

	line = strings.Trim(line, `"`)

	return fmt.Sprintf(format, color.Apply(line, stringColor))
}

// colorKey colorizes yaml keys
func (yp *YamlPrinter) colorKey(key string) string {
	format := "%s"
	if strings.HasSuffix(key, ":") {
		format += ":"
	}
	if yp.inArray(key) {
		format = "- " + format
	}

	key = strings.TrimSuffix(key, ":")
	key = strings.TrimPrefix(key, "- ")

	return fmt.Sprintf(format, color.Apply(key, keyColor))
}

// colorValue colorizes yaml values
func (yp *YamlPrinter) colorValue(value string) string {
	if value == "{}" {
		return "{}"
	}

	format := "%s"
	if yp.isString(value) {
		format = `"%s"`
	}
	if yp.inArray(value) {
		format = "- " + format
	}

	value = strings.TrimPrefix(value, "- ")
	value = strings.Trim(value, `"`)

	return fmt.Sprintf(format, color.Apply(value, getColorByValueType(value)))
}
