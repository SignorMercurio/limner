package transformer

import (
	"fmt"

	"github.com/SignorMercurio/limner/color"
)

type FormatTransformer struct {
	InType  string
	OutType string
}

func (ft *FormatTransformer) Transform(input []byte) ([]byte, error) {
	var trans Transformer = &PlainTransformer{}

	switch ft.OutType {
	case "yaml", "yml":
		trans = NewYamlTransformer(ft.InType)
	case "json":
		trans = NewJsonTransformer(ft.InType)
	}

	return trans.Transform(input)
}

func unknownInput() {
	fmt.Println(color.Apply("Unknown input format, using default transformer", color.Yellow))
}
