package printer

import (
	"io"
)

type ColorPrinter struct {
	Type string
	Args []string
}

func (cp *ColorPrinter) Print(r io.Reader, w io.Writer) {
	var printer Printer = &SingleColorPrinter{Color: StringColor}

	// change the printer if type is already enforced
	switch cp.Type {
	case "yaml", "yml":
		printer = &YamlPrinter{}
		// case "json":
		// case "table":
		// case "xml":
	}

	// otherwise, try to determine the type
	// s := bufio.NewScanner(r)
	// if shouldYaml(cp.Args, s) {
	// 	printer = &YamlPrinter{}
	// } else if shouldJson(cp.Args, s) {
	// 	printer = &JsonPrinter{}
	// } else if shouldTable(cp.Args, s) {
	// 	printer = &TablePrinter{}
	// }

	// Finally, we can print something
	printer.Print(r, w)
}
