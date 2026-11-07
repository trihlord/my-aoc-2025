package day4

import (
	"github.com/trihlord/myaoc2025/internal/file"
	"github.com/trihlord/myaoc2025/internal/rolls"
)

var byteOffsets = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func readRolls(filename string) ([]rolls.Roll, int, error) {
	j, l := 0, 0
	rr := []rolls.Roll{}
	for line, err := range file.ReadLines(filename) {
		if err != nil {
			return nil, 0, err
		}
		if j == 0 {
			l = len(line)
		}
		for _, b := range []byte(line) {
			rr = append(rr, rolls.Roll(b))
		}
		j++
	}
	return rr, l, nil
}

func isForkliftableRoll(rr []rolls.Roll, l, i int) bool {
	c, x, y, mx, my := 0, i%l, i/l, l, len(rr)/l
	if !rr[i].IsForkliftable() {
		return false
	}
	for _, o := range byteOffsets {
		if ox, oy := x+o[0], y+o[1]; oy >= 0 && oy < my && ox >= 0 && ox < mx && rr[oy*l+ox].IsForkliftable() {
			c++
		}
	}
	return c < 4
}

func CountRollsForkliftable(filename string) (int, error) {
	c := 0
	rr, l, err := readRolls(filename)
	if err != nil {
		return 0, err
	}
	for i := range rr {
		if isForkliftableRoll(rr, l, i) {
			c++
		}
	}
	return c, nil
}
