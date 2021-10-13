package printer

import (
	"io"
	"strconv"
	"strings"

	"github.com/SignorMercurio/limner/color"
	"gopkg.in/yaml.v3"
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

func shouldYaml(buf string) bool {
	// Look at the first 3 lines, so split into 3+1 parts
	lines := 3
	splitted := strings.SplitN(buf, "\n", lines+1)
	// in case of output shorter than 3 lines
	actualLines := len(splitted)
	if actualLines < lines {
		lines = actualLines
	}

	// Prepare for yaml unmarshal
	result := ""
	for i := 0; i < lines; i++ {
		result += splitted[i] + "\n"
	}

	out := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(result), out)

	return err == nil
}

func shouldTable(buf string) bool {
	lines := 2
	splitted := strings.SplitN(buf, "\n", lines+1)
	actualLines := len(splitted)
	if actualLines < lines {
		lines = actualLines
	}

	for i := 0; i < lines; i++ {
		if splitted[i] != "" && strings.ToUpper(splitted[i]) == splitted[i] {
			return true
		}
	}
	return false
}

func seemNegative(status string) bool {
	negativeKeywords := []string{
		"fail",
		"backoff",
		"exceed",
		"not",
		"err",
		"invalid",
		"unable",
		"unhealthy",
		"unknown",
		"evict",
		"bad",
		"timeout",
		"panic",
		"fatal",
	}

	for _, v := range negativeKeywords {
		if strings.Contains(status, v) {
			return true
		}
	}
	return false
}

func seemWarning(status string) bool {
	return strings.Contains(status, "ing")
}

func seemPositive(status string) bool {
	positiveKeywords := []string{
		"ok",
		"ted",
		"led",
		"ged",
		"zed",
		"success",
		"succeed",
		"ready",
		"normal",
		"healthy",
		"running",
		"done",
	}

	for _, v := range positiveKeywords {
		if strings.Contains(status, v) {
			return true
		}
	}
	return false
}

func ColorStatus(_ int, status string) (color.Color, bool) {
	status = strings.ToLower(status)
	switch {
	// the order matters!
	case seemNegative(status):
		return color.Red, true
	case seemPositive(status):
		return color.Green, true
	case seemWarning(status):
		return color.Yellow, true
	case strings.Count(status, "/") == 1:
		ready := strings.Split(status, "/")
		if ready[0] == ready[1] {
			return color.Green, true
		} else {
			_, e1 := strconv.Atoi(ready[0])
			_, e2 := strconv.Atoi(ready[1])
			if e1 == nil && e2 == nil {
				return color.Yellow, true
			}
		}
	}
	return 0, false
}
