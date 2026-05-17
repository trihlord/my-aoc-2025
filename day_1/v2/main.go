package main

import "my-aoc-2025/internal/rotation"

func Password(ss []string) (int, error) {
	p, d := 0, 50

	for _, s := range ss {
		r, err := rotation.Parse(s)
		if err != nil {
			return 0, err
		}

		var tp, td int
		if td = d + r; td < 0 {
			if d != 0 {
				tp++
			}
			td *= -1
		} else if td == 0 {
			tp++
		}
		tp += td / 100

		// TODO: td == 0 ?
		d = ((d+r)%100 + 100) % 100

		p += tp
	}

	return p, nil
}
