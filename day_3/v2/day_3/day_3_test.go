package day3

import (
	"path/filepath"
	"testing"
)

func TestSumMaximumTerm(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "first",
			in:   "987654321111111",
			out:  98,
		},
		{
			name: "second",
			in:   "811111111111119",
			out:  89,
		},
		{
			name: "third",
			in:   "234234234234278",
			out:  78,
		},
		{
			name: "fourth",
			in:   "818181911112111",
			out:  92,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			term, err := SumMaximumTerm(tC.in)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := term, tC.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}

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
			out:  357,
		},
		{
			name: "input",
			in:   "input.txt",
			out:  17343,
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
