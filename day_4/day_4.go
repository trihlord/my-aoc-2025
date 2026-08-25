package day4

import (
	"errors"

	"github.com/trihlord/myaoc2025/internal/file"
)

const (
	byteDot byte = '.'
	byteAt  byte = '@'
)

var errInvalidByte = errors.New("invalid byte")

func isInvalidByte(b byte) bool {
	return b != byteDot && b != byteAt
}

func isCountableByte(bb []byte, d int, x int, y int) bool {
	c := 0
	if bb[len(bb)-1] == byteDot {
		return false
	}
	if y > 1 && x > 1 && bb[d*(y-2)+x-2] == byteAt {
		c++
	}
	if y > 1 && x > 0 && bb[d*(y-2)+x-1] == byteAt {
		c++
	}
	if y > 1 && bb[d*(y-2)+x] == byteAt {
		c++
	}
	if y > 0 && x > 1 && bb[d*(y-1)+x-2] == byteAt {
		c++
	}
	if y > 0 && x > 0 && bb[d*(y-1)+x-1] == byteAt {
		c++
	}
	if y > 0 && bb[d*(y-1)+x] == byteAt {
		c++
	}
	if x > 1 && bb[d*y+x-2] == byteAt {
		c++
	}
	if x > 0 && bb[d*y+x-1] == byteAt {
		c++
	}
	return c < 4
}

func CountRollsForkliftable(filename string) (int, error) {
	bb := []byte{}
	c, i := 0, 0
	for line, err := range file.ReadLines(filename) {
		if err != nil {
			return 0, nil
		}
		for j, b := range []byte(line) {
			if isInvalidByte(b) {
				return 0, errInvalidByte
			}
			if bb = append(bb, b); isCountableByte(bb, len(line), j, i) {
				c++
			}
		}
		i++
	}
	return c, nil
}
