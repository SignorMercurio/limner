package color

import (
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestApply(t *testing.T) {
	val := "hello"
	dst := "\x1b[31mhello\x1b[0m"

	testutil.MustEqual(t, dst, Apply(val, Red))
}
