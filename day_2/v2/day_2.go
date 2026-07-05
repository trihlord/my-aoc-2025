package day2

import (
	"errors"
	"strconv"
	"strings"

	"github.com/trihlord/myaoc2025/internal/digits"
	"github.com/trihlord/myaoc2025/internal/file"
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

func SumInvalidIdRanges(fileName string) (int, error) {
	sum := 0
	for line, err := range file.ReadCommas(fileName) {
		if err != nil {
			return 0, err
		}
		idRange, err := newIdRange(line)
		if err != nil {
			return 0, err
		}
		for id := idRange.begin; id <= idRange.end; id++ {
			if digits.IsRepeated(id) {
				sum += id
			}
		}
	}
	return sum, nil
}
