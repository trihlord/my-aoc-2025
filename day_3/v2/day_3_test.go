package day3

import (
	"path/filepath"
	"testing"
)

func TestSumMaximum(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "example",
			in:   "example.txt",
			out:  3121910778619,
		},
		{
			name: "input",
			in:   "input.txt",
			out:  172664333119298, // TODO: calculate
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			sum, err := SumMaximum(filepath.Join("testdata", tC.in))
			if err != nil {
				t.Fatal(err)
			}
			if got, want := sum, tC.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
