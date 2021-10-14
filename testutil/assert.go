package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func MustEqual(t testing.TB, expected, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("diff (-expected +got):\n%s", diff)
	}
}
