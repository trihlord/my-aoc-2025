package day2

import (
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		in   string
		out  int
	}{
		{
			desc: "example",
			in:   "example.txt",
			out:  1227775554,
		},
		{
			desc: "input",
			in:   "input.txt",
			out:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			sum, err := SumInvalidIdRanges(filepath.Join("testdata", tC.in))
			if err != nil {
				t.Fatal(err)
			}
			if sum != tC.out {
				t.Fatalf("got %d, want %d", sum, tC.out)
			}
		})
	}
}
