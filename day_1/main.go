package main

import (
	"errors"
	"strconv"
)

func Password(rotations []string) (int, error) {
	count, point := 0, 50

	for _, r := range rotations {
		d := r[0:1]
		if d != "L" && d != "R" {
			return 0, errors.New("rotation direction must be L or R")
		}

		c, err := strconv.Atoi(r[1:])
		if err != nil {
			return 0, errors.New("rotation clics must be positive integer")
		}

		if d == "L" {
			point -= c % 100
		}

		if d == "R" {
			point += c % 100
		}

		if point < 0 {
			point += 100
		}

		if point > 99 {
			point -= 100
		}

		if point == 0 {
			count += 1
		}
	}

	return count, nil
}
