package sequence

import "testing"

func TestMax(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		in   struct {
			str string
			cap int
		}
		out string
	}{
		{
			name: "starting left in order of two",
			in: struct {
				str string
				cap int
			}{
				str: "987654321111111",
				cap: 2,
			},
			out: "98",
		},
		{
			name: "starting left in order of twelve",
			in: struct {
				str string
				cap int
			}{
				str: "987654321111111",
				cap: 12,
			},
			out: "987654321111",
		},
		{
			name: "starting left out of order of two",
			in: struct {
				str string
				cap int
			}{
				str: "811111111111119",
				cap: 2,
			},
			out: "89",
		},
		{
			name: "starting left out of order of twelve",
			in: struct {
				str string
				cap int
			}{
				str: "811111111111119",
				cap: 12,
			},
			out: "811111111119",
		},
		{
			name: "starting middle out of order of two",
			in: struct {
				str string
				cap int
			}{
				str: "234234234234278",
				cap: 2,
			},
			out: "78",
		},
		{
			name: "starting middle out of order of twelve",
			in: struct {
				str string
				cap int
			}{
				str: "234234234234278",
				cap: 12,
			},
			out: "434234234278",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			t.Parallel()
			if got, want := Max(tC.in.str, tC.in.cap), tC.out; got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}
