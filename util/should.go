package util

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

type Obj = map[string]interface{}

// ShouldYaml returns if the buf should be in yaml format
func ShouldYaml(buf string) bool {
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

	var (
		outObj      Obj
		outArray    []Obj
		resultBytes = []byte(result)
	)
	errObj := yaml.Unmarshal(resultBytes, &outObj)
	errArray := yaml.Unmarshal(resultBytes, &outArray)

	return (errObj == nil) || (errArray == nil)
}

func ShouldJson(buf []byte) (interface{}, bool) {
	if jsonObj, ok := ShouldJsonObj(buf); ok {
		return jsonObj, true
	} else if jsonArray, ok := ShouldJsonArray(buf); ok {
		return *jsonArray, true
	}
	return nil, false
}

// ShouldJsonObj returns if the buf should be in simple json format
func ShouldJsonObj(buf []byte) (*Obj, bool) {
	var jsonObj Obj
	err := json.Unmarshal(buf, &jsonObj)
	return &jsonObj, err == nil
}

// ShouldJsonArray returns if the buf should be in json array format
func ShouldJsonArray(buf []byte) (*[]Obj, bool) {
	var jsonArray []Obj
	err := json.Unmarshal(buf, &jsonArray)
	return &jsonArray, err == nil
}

// ShouldTable returns if the buf should be in table format
func ShouldTable(buf string) bool {
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
