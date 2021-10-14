package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

type ColorPrinter struct {
	Type string
}

func (cp *ColorPrinter) Print(buf string, w io.Writer) {
	var printer Printer = &SingleColorPrinter{Color: stringColor}
	var jsonObj map[string]interface{}

	// change the printer if type is already enforced
	switch cp.Type {
	case "yaml", "yml":
		printer = NewYamlPrinter()
	case "json":
		if err := json.Unmarshal([]byte(buf), &jsonObj); err != nil {
			fmt.Fprintln(w, color.Apply("Failed to unmarshal json, using default printer", color.Yellow))
		} else {
			printer = NewJsonPrinter(jsonObj)
		}
	case "table":
		printer = NewTablePrinter(ColorStatus)
		// case "xml":
	default:
		// otherwise, try to determine the type
		switch {
		case shouldJson(buf, &jsonObj):
			printer = NewJsonPrinter(jsonObj)
		case shouldTable(buf):
			printer = NewTablePrinter(ColorStatus)
		case shouldYaml(buf):
			printer = NewYamlPrinter()
		}
	}
	// Finally, we can print something
	printer.Print(buf, w)
}

// ColorStatus colorizes the status based on status texts
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
