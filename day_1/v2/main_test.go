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
			name: "example",
			in:   "example.txt",
			out:  6,
		},
		{
			name: "input",
			in:   "input.txt",
			out:  6913,
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
