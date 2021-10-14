package printer

import (
	"strconv"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

// getIndent returns the length of indent in spaces for a line
func getIndent(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

// toSpaces returns n spaces
func toSpaces(cnt int) string {
	return strings.Repeat(" ", cnt)
}

// isNull returns if the value is sort of null value
func isNull(value string) bool {
	nulls := []string{"null", "<none>", "<unknown>", "<unset>", "<nil>"}
	for _, v := range nulls {
		if value == v {
			return true
		}
	}
	return false
}

// isBool returns if the value is of type bool
func isBool(value string) bool {
	return strings.ToLower(value) == "true" || strings.ToLower(value) == "false"
}

// isNum returns if the value can be transformed into a number
func isNum(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

// getColorByValueType returns a color to use based on the type of the value
func getColorByValueType(value string) color.Color {
	switch {
	case isNull(value):
		return nullColor
	case isBool(value):
		return boolColor
	case isNum(value):
		return numberColor
	default:
		return stringColor
	}
}
