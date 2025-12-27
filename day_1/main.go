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
			point = point - c
		}

		if d == "R" {
			point = point + c
		}

		m := point % 100

		if m < 0 {
			point = m + 100
		} else {
			point = m
		}

		if point == 0 {
			count += 1
		}
	}

	return count, nil
}

func PasswordV2(rotations []string) (int, error) {
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
			point -= c
		}

		if d == "R" {
			point += c
		}

		q, m := point/100, point%100

		if m < 0 {
			point = m + 100
		} else {
			point = m
		}

		if q < 0 {
			count -= q
		} else {
			count += q
		}

		if m < 0 && 0 < point {
			count++
		}

		if point == 0 && 0 < m {
			count++
		}
	}

	return count, nil
}
