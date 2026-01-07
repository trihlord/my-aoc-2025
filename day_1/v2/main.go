package main

import (
	"errors"
	"strconv"
)

func Password(ss []string) (int, error) {
	p, d := 0, 50

	for _, s := range ss {
		r, err := parseRotation(s)
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

func parseRotation(s string) (int, error) {
	d := s[0:1]
	if d != "L" && d != "R" {
		return 0, errors.New("rotation direction must be L or R")
	}

	r, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, errors.New("rotation clics must be positive integer")
	}

	if d == "L" {
		r *= -1
	}

	return r, nil
}
