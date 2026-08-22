package day3

import (
	"strconv"

	"github.com/trihlord/myaoc2025/internal/file"
	"github.com/trihlord/myaoc2025/internal/sequence"
)

func SumMaximum(fileName string) (int, error) {
	sum := 0
	for line, err := range file.ReadLines(fileName) {
		if err != nil {
			return 0, err
		}
		term, err := strconv.Atoi(sequence.Max(line, 2))
		if err != nil {
			return 0, err
		}
		sum += term
	}
	return sum, nil
}
