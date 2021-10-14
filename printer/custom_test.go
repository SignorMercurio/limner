package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/color"
	"github.com/SignorMercurio/limner/testutil"
)

func TestCustomPrinter(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "echo -n 'hello'",
			src:  `hello`,
			dst: `[32mhello[0m
`,
		},
		{
			name: "echo -n 'world'",
			src:  `world`,
			dst: `[34mworld[0m
`,
		},
		{
			name: "echo -n 'foobar'",
			src:  `foobar`,
			dst: `[37mfoobar[0m
`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			p := &CustomPrinter{
				ColorPicker: func(line string) color.Color {
					if line == "hello" {
						return color.Green
					} else if line == "world" {
						return color.Blue
					}
					return color.White
				},
			}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
