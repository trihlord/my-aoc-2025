package rotation

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string
		out  int
	}{
		{name: "left", in: "L50", out: -50},
		{name: "right", in: "R50", out: 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			r, err := Parse(test.in)
			if err != nil {
				t.Fatal(err)
			}

			if got, want := r, test.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
