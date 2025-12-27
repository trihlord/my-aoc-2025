package main

import (
	"errors"
	"strconv"
)

func Password(ss []string) (int, error) {
	zeroes, dial := 0, 50

	for _, s := range ss {
		r, err := rotation(s)
		if err != nil {
			return 0, err
		}

		if dial = ((dial+r)%100 + 100) % 100; dial == 0 {
			zeroes++
		}
	}

	return zeroes, nil
}

func rotation(s string) (int, error) {
	d := s[0:1]
	if d != "L" && d != "R" {
		return 0, errors.New("rotation direction must be L or R")
	}

	c, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, errors.New("rotation clics must be positive integer")
	}

	if d == "L" {
		c *= -1
	}

	return c, nil
}
