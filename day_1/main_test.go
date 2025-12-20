package main

import (
	"my-aoc-2025/aoc"
	"path/filepath"
	"testing"
)

func TestPassword(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		t.Parallel()

		lines, err := aoc.ReadLines(filepath.Join("testdata", "example.txt"))
		if err != nil {
			t.Fatal(err)
		}

		password, err := Password(lines)
		if err != nil {
			t.Fatal(err)
		}

		if password != 3 {
			t.Fatalf("got %d, want %d", password, 3)
		}
	})

	t.Run("Input", func(t *testing.T) {
		t.Parallel()

		lines, err := aoc.ReadLines(filepath.Join("testdata", "input.txt"))
		if err != nil {
			t.Fatal(err)
		}

		password, err := Password(lines)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("got %d", password)
	})
}
