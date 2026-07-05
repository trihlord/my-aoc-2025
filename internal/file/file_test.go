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
			name: filepath.Join("testdata", "empty.txt"),
			line: "",
		},
		{
			desc: "returns single line",
			name: filepath.Join("testdata", "single_line.txt"),
			line: "single",
		},
		{
			desc: "returns joined two lines",
			name: filepath.Join("testdata", "two_lines.txt"),
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
