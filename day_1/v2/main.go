package main

import "my-aoc-2025/internal/rotation"

func Password(ss []string) (int, error) {
	p, d := 0, 50

	for _, s := range ss {
		r, err := rotation.Parse(s)
		if err != nil {
			return 0, err
		}

		w := (d + r) / 100
		if w < 0 {
			w *= -1
		}

		d = ((d+r)%100 + 100) % 100

		p += w
	}

	return p, nil
}
