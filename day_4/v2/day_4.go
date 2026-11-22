package day4

import (
	"github.com/trihlord/myaoc2025/internal/file"
	"github.com/trihlord/myaoc2025/internal/rolls"
)

func newRollWarehouse(filename string) (*rolls.RollWarehouse, error) {
	var rw *rolls.RollWarehouse
	for line, err := range file.ReadLines(filename) {
		if err != nil {
			return nil, err
		}
		if rw == nil {
			rw = rolls.NewRollWarehouse(len(line))
		}
		for _, b := range []byte(line) {
			rw.Append(rolls.Roll(b))
		}
	}
	return rw, nil
}

func CountRollsForkliftable(filename string) (int, error) {
	rw, err := newRollWarehouse(filename)
	if err != nil {
		return 0, err
	}
	c := 0
	for i := range rw.Len() {
		if rw.IsForkliftableAt(i) {
			c++
		}
	}
	return c, nil
}
