package transformer

import (
	"github.com/SignorMercurio/limner/util"
	"gopkg.in/yaml.v3"
)

type YamlTransformer struct {
	Type string
	obj  interface{}
}

func (yt *YamlTransformer) Transform(input []byte) ([]byte, error) {
	var isJson bool
	// fill the obj if the type is already enforced
	switch yt.Type {
	case "json":
		if yt.obj, isJson = util.ShouldJson(input); isJson {
			return yaml.Marshal(yt.obj)
		}
	// case "...":
	default:
		// otherwise, try to determine the type
		if yt.obj, isJson = util.ShouldJson(input); isJson {
			return yaml.Marshal(yt.obj)
		}
	}

	unknownInput()
	return input, nil
}

func NewYamlTransformer(Type string) *YamlTransformer {
	return &YamlTransformer{
		Type: Type,
	}
}
