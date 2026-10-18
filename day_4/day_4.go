package day4

import (
	"errors"

	"github.com/trihlord/myaoc2025/internal/file"
)

const (
	byteDot byte = '.'
	byteAt  byte = '@'
)

var (
	errInvalidByte = errors.New("invalid byte")
	byteOffsets    = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

func isInvalidByte(b byte) bool {
	return b != byteDot && b != byteAt
}

func readRollBytes(filename string) ([]byte, int, error) {
	j, l := 0, 0
	bb := []byte{}
	for line, err := range file.ReadLines(filename) {
		if err != nil {
			return nil, 0, err
		}
		if j == 0 {
			l = len(line)
		}
		for _, b := range []byte(line) {
			if isInvalidByte(b) {
				return nil, 0, errInvalidByte
			}
			bb = append(bb, b)
		}
		j++
	}
	return bb, l, nil
}

func isForkliftableByte(bb []byte, l, i int) bool {
	c, x, y, mx, my := 0, i%l, i/l, l, len(bb)/l
	if bb[i] != byteAt {
		return false
	}
	for _, o := range byteOffsets {
		if ox, oy := x+o[0], y+o[1]; oy >= 0 && oy < my && ox >= 0 && ox < mx && bb[oy*l+ox] == byteAt {
			c++
		}
	}
	return c < 4
}

func CountRollsForkliftable(filename string) (int, error) {
	c, l := 0, 0
	bb, l, err := readRollBytes(filename)
	if err != nil {
		return 0, err
	}
	for i := range bb {
		if isForkliftableByte(bb, l, i) {
			c++
		}
	}
	return c, nil
}
