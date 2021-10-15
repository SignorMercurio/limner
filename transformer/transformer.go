package transformer

type Transformer interface {
	Transform(input []byte) ([]byte, error)
}
