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
			name: "is repeated single digit twice",
			in:   11,
			out:  true,
		},
		{
			name: "is repeated single digit thrice",
			in:   111,
			out:  true,
		},
		{
			name: "is repeated two digits twice",
			in:   1010,
			out:  true,
		},
		{
			name: "is repeated five digits twice",
			in:   1188511885,
			out:  true,
		},
		{
			name: "is repeated three digits twice",
			in:   222222,
			out:  true,
		},
		{
			name: "is repeated four digits twice",
			in:   38593859,
			out:  true,
		},
		{
			name: "is repeated two digits thrice",
			in:   565656,
			out:  true,
		},
		{
			name: "is repeated tree digits thrice",
			in:   824824824,
			out:  true,
		},
		{
			name: "is repeated two digits fives",
			in:   2121212121,
			out:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			if got, want := IsRepeated(tC.in), tC.out; got != want {
				t.Fatalf("got %t, want %t", got, want)
			}
		})
	}
}
