package transformer

import (
	"encoding/json"

	"github.com/SignorMercurio/limner/util"
	"gopkg.in/yaml.v3"
)

type JsonTransformer struct {
	Type string
	obj  util.Obj
}

func (jt *JsonTransformer) Transform(input []byte) ([]byte, error) {
	// fill the obj if the type is already enforced
	switch jt.Type {
	case "yaml":
		if err := yaml.Unmarshal(input, &jt.obj); err == nil {
			return json.MarshalIndent(jt.obj, "", "    ")
		}
	default:
		// otherwise, try to determine the type
		switch {
		case util.ShouldYaml(string(input)):
			if err := yaml.Unmarshal(input, &jt.obj); err == nil {
				return json.MarshalIndent(jt.obj, "", "    ")
			}
		}
	}

	unknownInput()
	return input, nil
}

func NewJsonTransformer(Type string) *JsonTransformer {
	return &JsonTransformer{
		Type: Type,
	}
}
