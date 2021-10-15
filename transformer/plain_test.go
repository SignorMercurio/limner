package transformer

import (
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestPlainTransformer(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "simple case",
			src:  `hello`,
			dst:  `hello`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			trans := &FormatTransformer{}
			w, _ := trans.Transform([]byte(tt.src))
			testutil.MustEqual(t, []byte(tt.dst), w)
		})
	}
}
