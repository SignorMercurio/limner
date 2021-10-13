package printer

import (
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

type ColorPrinter struct {
	Type string
}

func (cp *ColorPrinter) Print(buf string, w io.Writer) {
	var printer Printer = &SingleColorPrinter{Color: StringColor}

	// change the printer if type is already enforced
	switch cp.Type {
	case "yaml", "yml":
		printer = NewYamlPrinter()
		// case "json":
	case "table":
		printer = NewTablePrinter(ColorStatus)
		// case "xml":
	default:
		// otherwise, try to determine the type
		switch {
		case shouldYaml(buf):
			printer = NewYamlPrinter()
		// case shouldJson(buf):
		// 	printer = NewJsonPrinter()
		case shouldTable(buf):
			printer = NewTablePrinter(ColorStatus)
		}
	}
	// Finally, we can print something
	printer.Print(buf, w)
}

func ColorStatus(_ int, status string) (color.Color, bool) {
	status = strings.ToLower(status)
	switch {
	// the order matters!
	case seemsNegative(status):
		return color.Red, true
	case seemsPositive(status):
		return color.Green, true
	case seemsWarning(status):
		return color.Yellow, true
	default:
		return seemsReadyStatus(status)
	}
}
