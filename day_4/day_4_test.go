package day4

import (
	"path/filepath"
	"testing"
)

func TestCountRollsForkliftable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "example",
			in:   "example.txt",
			out:  13,
		},
		// TODO: find valid output
		// {
		// 	name: "input",
		// 	in:   "input.txt",
		// 	out:  0,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			i, err := CountRollsForkliftable(filepath.Join("testdata", tC.in))
			if err != nil {
				t.Fatal(err)
			}
			if got, want := i, tC.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
