package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/SignorMercurio/limner/color"
)

type JsonPrinter struct {
	jsonObj interface{}
}

func (jp *JsonPrinter) Print(buf string, w io.Writer) {
	jsonBytes, _ := json.MarshalIndent(jp.jsonObj, "", "    ")

	for _, line := range strings.Split(string(jsonBytes), "\n") {
		jp.printJson(line, w)
	}
}

func NewJsonPrinter(jsonObj interface{}) *JsonPrinter {
	return &JsonPrinter{
		jsonObj: jsonObj,
	}
}

// printJson prints a line to writer w in json format
func (jp *JsonPrinter) printJson(line string, w io.Writer) {
	indentCnt := getIndent(line)
	indent := toSpaces(indentCnt)
	line = strings.TrimLeft(line, " ")

	// { OR }(,) OR ](,)
	if strings.HasPrefix(line, "{") || strings.HasPrefix(line, "}") || strings.HasPrefix(line, "]") {
		fmt.Fprintf(w, "%s%s\n", indent, line)
		return
	}

	// "key": { OR "key": [ OR "key": value(,) OR value(,)
	splitted := strings.SplitN(line, ": ", 2)
	key := splitted[0]

	// value(,)
	if len(splitted) == 1 {
		fmt.Fprintf(w, "%s%s\n", indent, jp.colorValue(key))
		return
	}

	value := splitted[1]
	fmt.Fprintf(w, "%s%s: %s\n", indent, jp.colorKey(key), jp.colorValue(value))
}

// colorKey colorizes json keys
func (jp *JsonPrinter) colorKey(key string) string {
	key = strings.TrimSuffix(key, ":")
	key = strings.Trim(key, `"`)

	return fmt.Sprint(color.Apply(key, keyColor))
}

// colorValue colorizes json values
func (jp *JsonPrinter) colorValue(value string) string {
	switch value {
	case "{", "[", "{}", "{},":
		return value
	}

	format := "%s"
	if jp.isString(value) {
		format = `"%s"`
	}
	if strings.HasSuffix(value, ",") {
		format += ","
	}

	value = strings.TrimRight(value, ",")
	value = strings.Trim(value, `"`)

	return fmt.Sprintf(format, color.Apply(value, getColorByValueType(value)))
}

// isString returns if the line is double quoted
func (jp *JsonPrinter) isString(line string) bool {
	return strings.HasPrefix(line, `"`) && (strings.HasSuffix(line, `"`) || strings.HasSuffix(line, `",`))
}
