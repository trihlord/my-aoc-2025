package digits

import "testing"

func TestIsRepeated(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		in   int
		out  bool
	}{
		{
			name: "",
			in:   11,
			out:  true,
		},
		{
			name: "",
			in:   111,
			out:  true,
		},
		{
			name: "",
			in:   1188511885,
			out:  true,
		},
		{
			name: "",
			in:   2121212121,
			out:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			if got, want := IsRepeated(tC.in), tC.out; got != want {
				t.Fatalf("got %d, want %d", got, want)
			}
		})
	}
}
