package day4

import (
	"github.com/trihlord/myaoc2025/internal/file"
	"github.com/trihlord/myaoc2025/internal/rolls"
)

func newRollNeighbours(filename string) (rolls.RollNeighbours, int, error) {
	j, l := 0, 0
	rns := rolls.RollNeighbours{}
	for line, err := range file.ReadLines(filename) {
		if err != nil {
			return nil, 0, err
		}
		if j == 0 {
			l = len(line)
		}
		for _, b := range []byte(line) {
			rns.Append(rolls.Roll(b))
		}
		j++
	}
	return rns, l, nil
}

func CountRollsForkliftable(filename string) (int, error) {
	c := 0
	rns, l, err := newRollNeighbours(filename)
	if err != nil {
		return 0, err
	}
	for i := range rns {
		if rns.IsForkliftableAt(l, i) {
			c++
		}
	}
	return c, nil
}
