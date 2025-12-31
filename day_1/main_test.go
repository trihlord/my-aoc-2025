package main

import (
	"my-aoc-2025/aoc"
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
			out:  3,
		},
		{
			name: "input",
			in:   "input.txt",
			out:  1182,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			lines, err := aoc.ReadLines(filepath.Join("testdata", test.in))
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

func TestPasswordV2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "example short",
			in:   "example_short.txt",
			out:  2,
		},
		{
			name: "example shorter",
			in:   "example_shorter.txt",
			out:  5,
		},
		{
			name: "example",
			in:   "example.txt",
			out:  6,
		},
		{
			name: "input",
			in:   "input.txt",
			out:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			lines, err := aoc.ReadLines(filepath.Join("testdata", test.in))
			if err != nil {
				t.Fatal(err)
			}

			password, err := PasswordV2(lines)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := password, test.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
