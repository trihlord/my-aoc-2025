package day3

import (
	"bufio"
	"iter"
	"os"
	"strconv"
)

func readLines(fileName string) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		file, err := os.Open(fileName)
		if err != nil {
			yield("", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}

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
	for line, err := range readLines(fileName) {
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
