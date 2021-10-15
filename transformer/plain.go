package transformer

type PlainTransformer struct{}

func (pt *PlainTransformer) Transform(input []byte) ([]byte, error) {
	return input, nil
}
