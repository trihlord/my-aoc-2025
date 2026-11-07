package rolls

import "testing"

func TestIsForkliftable(t *testing.T) {
	testCases := []struct {
		name string
		in   byte
		out  bool
	}{
		{
			name: "at byte",
			in:   '@',
			out:  true,
		},
		{
			name: "dot byte",
			in:   '.',
			out:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			if got, want := Roll(tC.in).IsForkliftable(), tC.out; got != want {
				t.Fatalf("got %t, want %t", got, want)
			}
		})
	}
}
