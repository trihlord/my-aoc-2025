package day2

import (
	"path/filepath"
	"testing"
)

func TestSumInvalidIdRanges(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		in   string
		out  int
	}{
		{
			desc: "example",
			in:   "example.txt",
			out:  4174379265,
		},
		{
			desc: "input",
			in:   "input.txt",
			out:  54486209192,
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
