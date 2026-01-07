package rotation

import (
	"errors"
	"strconv"
)

var (
	errInvalidDirection = errors.New("rotation direction must be L or R")
	errInvalidClicks    = errors.New("rotation clics must be positive integer")
)

func Parse(s string) (int, error) {
	d := s[0:1]
	if d != "L" && d != "R" {
		return 0, errInvalidDirection
	}

	r, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, errInvalidClicks
	}

	if d == "L" {
		r *= -1
	}

	return r, nil
}
