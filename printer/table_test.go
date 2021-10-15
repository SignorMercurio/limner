package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestTablePrinter(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "k get po",
			src: `NAME                     READY   STATUS              RESTARTS   AGE
nginx-7848d4b86f-2pq9t   1/1     Running             1          23h
nginx-7848d4b86f-mk9pw   0/1     CrashLoopBackoff    1          31h
nginx-7848d4b86f-tffzh   0/1     ContainerCreating   1          23h
`,
			dst: `[34mNAME                     READY   STATUS              RESTARTS   AGE[0m
[37mnginx-7848d4b86f-2pq9t[0m   [32m1/1[0m     [32mRunning[0m             [36m1[0m          [37m23h[0m
[37mnginx-7848d4b86f-mk9pw[0m   [33m0/1[0m     [31mCrashLoopBackoff[0m    [36m1[0m          [37m31h[0m
[37mnginx-7848d4b86f-tffzh[0m   [33m0/1[0m     [33mContainerCreating[0m   [36m1[0m          [37m23h[0m
[34m[0m
`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer
			p := &ColorPrinter{}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
			w.Reset()

			p = &ColorPrinter{Type: "table"}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
