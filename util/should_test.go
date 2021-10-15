package util

import (
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestShouldYaml(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  bool
	}{
		{
			name: "should be yaml",
			src: `a: 0
b: "1"
c:
  d:
  - eee
  - "fff"
  - 123`,
			dst: true,
		},
		{
			name: "array only should also be yaml",
			src: `- aaa: bbb
  bbb: "ccc"
- ccc: "ddd"
  ddd: eee`,
			dst: true,
		},
		{
			name: "should not be yaml",
			src: `a:0
b-c:'d'`,
			dst: false,
		},
		{
			name: "one line only",
			src:  `hello world!`,
			dst:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testutil.MustEqual(t, tt.dst, ShouldYaml(tt.src))
		})
	}
}

func TestShouldJson(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  bool
	}{
		{
			name: "should be json",
			src: `{
    "a": 0,
    "b": "1",
    "c": {
		"d": [
			"eee",
			"fff",
			123
		]
	}
}`,
			dst: true,
		},
		{
			name: "array only should also be json",
			src: `[
	{
		"aaa": "bbb",
		"bbb": "ccc"
	},
	{
		"ccc": "ddd",
		"ddd": "eee"
	}
]`,
			dst: true,
		},
		{
			name: "should not be json",
			src: `{
	a:0
	b-c:'d'
}`,
			dst: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			src := []byte(tt.src)
			_, isJson := ShouldJson(src)
			if !isJson {
				_, isJson = ShouldJsonArray(src)
			}
			testutil.MustEqual(t, tt.dst, isJson)
		})
	}
}

func TestShouldTable(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  bool
	}{
		{
			name: "should be table",
			src: `COLUMNA  COLUMNB  COLUMNC
aaa  bbb  ccc`,
			dst: true,
		},
		{
			name: "should also be table",
			src: `col_a  col_b  col_c
---  ---  ---
aaa  bbb  ccc`,
			dst: true,
		},
		{
			name: "should not be table",
			src: `ColumnA  ColumnB  ColumnC
aaa  bbb  ccc`,
			dst: false,
		},
		{
			name: "one line only",
			src:  `hello world!`,
			dst:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testutil.MustEqual(t, tt.dst, ShouldTable(tt.src))
		})
	}
}
