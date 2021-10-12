package printer

import (
	"io"

	"github.com/SignorMercurio/limner/color"
)

type ColorPrinter struct {
	Type    string
	LightBg bool
	Args    []string
}

func (cp *ColorPrinter) Print(r io.Reader, w io.Writer) {
	var printer Printer = &SingleColorPrinter{Color: color.Green}

	// change the printer if type is already enforced
	// switch cp.Type {
	// case "yaml","yml":
	// case "json":
	// case "xml":
	// case "table":
	// }

	// otherwise, try to determine the type
	// s := bufio.NewScanner(r)
	// if shouldYaml(cp.Args, s) {
	// 	printer = &YamlPrinter{
	// 		LightBg: cp.LightBg,
	// 	}
	// } else if shouldJson(cp.Args, s) {
	// 	printer = &JsonPrinter{
	// 		LightBg: cp.LightBg,
	// 	}
	// } else if shouldTable(cp.Args, s) {
	// 	printer = &TablePrinter{
	// 		LightBg: cp.LightBg,
	// 	}
	// }

	// Finally, we can print something
	printer.Print(r, w)
}
