package day2

import (
	"bufio"
	"bytes"
	"errors"
	"iter"
	"os"
	"strconv"
	"strings"
)

var errMultipleHyphens = errors.New("multiple hyphens")

type idRange struct {
	begin int
	end   int
}

func newIdRange(s string) (*idRange, error) {
	ss := strings.Split(s, "-")
	if len(ss) != 2 {
		return nil, errMultipleHyphens
	}
	begin, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(ss[1])
	if err != nil {
		return nil, err
	}
	return &idRange{begin, end}, nil
}

func scanIdRanges(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range data {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}
	return 0, nil, nil
}

func readIdRanges(fileName string) iter.Seq2[*idRange, error] {
	return func(yield func(*idRange, error) bool) {
		file, err := os.Open(fileName)
		if err != nil {
			yield(nil, err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(scanIdRanges)
		for scanner.Scan() {
			idRange, err := newIdRange(scanner.Text())
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(idRange, nil) {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			yield(nil, err)
		}
	}
}

func isInvalidId(id int) bool {
	s := strconv.Itoa(id)
	l := len(s)
	return s[l/2:] == s[:l/2]
}

func SumInvalidIdRanges(fileName string) (int, error) {
	sum := 0
	for idRange, err := range readIdRanges(fileName) {
		if err != nil {
			return 0, err
		}
		for id := idRange.begin; id <= idRange.end; id++ {
			if isInvalidId(id) {
				sum += id
			}
		}
	}
	return sum, nil
}
