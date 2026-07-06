package file

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		name string
		line string
	}{
		{
			desc: "returns empty line",
			name: filepath.Join("testdata", "split", "lines", "empty.txt"),
			line: "",
		},
		{
			desc: "returns one line",
			name: filepath.Join("testdata", "split", "lines", "one.txt"),
			line: "single",
		},
		{
			desc: "returns joined two lines",
			name: filepath.Join("testdata", "split", "lines", "two.txt"),
			line: "foobar",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lines := []string{}
			for line, err := range ReadLines(tC.name) {
				if err != nil {
					t.Fatal(err)
				}
				lines = append(lines, line)
			}
			if got, want := strings.Join(lines, ""), tC.line; got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestReadCommas(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		name string
		line string
	}{
		{
			desc: "returns empty line",
			name: filepath.Join("testdata", "split", "commas", "empty.txt"),
			line: "",
		},
		{
			desc: "returns one line",
			name: filepath.Join("testdata", "split", "commas", "none.txt"),
			line: "foo",
		},
		{
			desc: "returns joined one comma",
			name: filepath.Join("testdata", "split", "commas", "one.txt"),
			line: "foobar",
		},
		{
			desc: "returns joined two commas",
			name: filepath.Join("testdata", "split", "commas", "two.txt"),
			line: "foobarbaz",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			lines := []string{}
			for line, err := range ReadCommas(tC.name) {
				if err != nil {
					t.Fatal(err)
				}
				lines = append(lines, line)
			}
			if got, want := strings.Join(lines, ""), tC.line; got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}
