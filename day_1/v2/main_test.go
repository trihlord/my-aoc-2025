package main

import (
	"my-aoc-2025/internal/file"
	"path/filepath"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "example 1",
			in:   "example_1.txt",
			out:  1,
		},
		{
			name: "example 2",
			in:   "example_2.txt",
			out:  2,
		},
		{
			name: "example 3",
			in:   "example_3.txt",
			out:  3,
		},
		{
			name: "example 4",
			in:   "example_4.txt",
			out:  4,
		},
		{
			name: "example 5",
			in:   "example_5.txt",
			out:  5,
		},
		{
			name: "example 6",
			in:   "example_6.txt",
			out:  6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			lines, err := file.ReadLines(filepath.Join("testdata", test.in))
			if err != nil {
				t.Fatal(err)
			}

			password, err := Password(lines)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := password, test.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
