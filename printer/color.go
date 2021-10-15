package printer

import (
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
	"github.com/SignorMercurio/limner/util"
)

type ColorPrinter struct {
	Type string
}

func (cp *ColorPrinter) Print(buf string, w io.Writer) {
	var (
		printer Printer = &SingleColorPrinter{Color: stringColor}
		bufByte         = []byte(buf)
		jsonObj interface{}
		isJson  bool
	)

	// change the printer if type is already enforced
	switch cp.Type {
	case "yaml", "yml":
		printer = NewYamlPrinter()
	case "json":
		if jsonObj, isJson = util.ShouldJson(bufByte); isJson {
			printer = NewJsonPrinter(jsonObj)
		} else {
			fmt.Fprintln(w, color.Apply("Failed to unmarshal json, using default printer", color.Yellow))
		}
	case "table":
		printer = NewTablePrinter(ColorStatus)
	default:
		// otherwise, try to determine the type
		// try json first
		jsonObj, isJson = util.ShouldJson(bufByte)
		switch {
		case isJson:
			printer = NewJsonPrinter(jsonObj)
		case util.ShouldTable(buf):
			printer = NewTablePrinter(ColorStatus)
		case util.ShouldYaml(buf):
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
