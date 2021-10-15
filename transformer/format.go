package transformer

type FormatTransformer struct {
	InType  string
	OutType string
}

func (ft *FormatTransformer) Transform(input []byte) ([]byte, error) {
	var trans Transformer = &PlainTransformer{}

	switch ft.OutType {
	case "yaml", "yml":
		trans = NewYamlTransformer(ft.InType)
	}

	return trans.Transform(input)
}
