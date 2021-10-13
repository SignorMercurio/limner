package printer

import (
	"strings"

	"gopkg.in/yaml.v3"
)

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
