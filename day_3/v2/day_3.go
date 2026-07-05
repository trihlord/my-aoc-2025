package day3

import (
	"strconv"

	"github.com/trihlord/myaoc2025/internal/file"
)

func SumMaximumTerm(line string) (int, error) {
	dec, dig := 0, 0
	for _, r := range line {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			return 0, err
		}
		if i > dig {
			if dec > dig {
				dig = i
			} else {
				dec = dig
				dig = i
			}
		} else if i > dec {
			if dig > dec {
				dec = dig
				dig = i
			}
		} else if dig > dec {
			dec = dig
			dig = i
		}
	}
	return 10*dec + dig, nil
}

func SumMaximum(fileName string) (int, error) {
	sum := 0
	for line, err := range file.ReadLines(fileName) {
		if err != nil {
			return 0, err
		}
		term, err := SumMaximumTerm(line)
		if err != nil {
			return 0, err
		}
		sum += term
	}
	return sum, nil
}
