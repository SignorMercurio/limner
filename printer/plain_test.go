package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestPlainPrinter(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "echo -n 'hello'",
			src:  `hello`,
			dst:  `hello`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			p := &PlainPrinter{}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
