package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/color"
	"github.com/SignorMercurio/limner/testutil"
)

func TestSingleColorPrinter(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "echo -n 'hello'",
			src:  `hello`,
			dst: `[32mhello[0m`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			p := &SingleColorPrinter{color.Green}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
