package transformer

import (
	"encoding/json"
	"fmt"

	"github.com/SignorMercurio/limner/color"
	"github.com/SignorMercurio/limner/util"
	"gopkg.in/yaml.v3"
)

type YamlTransformer struct {
	Type string
	obj  map[string]interface{}
}

func (yt *YamlTransformer) Transform(input []byte) ([]byte, error) {
	// fill the obj if the type is already enforced
	switch yt.Type {
	case "json":
		if err := json.Unmarshal(input, &yt.obj); err != nil {
			return input, err
		} else {
			return yaml.Marshal(yt.obj)
		}
	default:
		// otherwise, try to determine the type
		switch {
		case util.ShouldJson(input, &yt.obj):
			return yaml.Marshal(yt.obj)
		}
	}

	fmt.Println(color.Apply("Unknown input format, using default transformer", color.Yellow))
	return input, nil
}

func NewYamlTransformer(Type string) *YamlTransformer {
	return &YamlTransformer{
		Type: Type,
	}
}
