package day2

import "testing"

func TestIds(t *testing.T) {
	t.Parallel()

	tests := []struct{ name string }{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Ids()
		})
	}
}
