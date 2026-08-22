package sequence

import "strconv"

func Max(s string, c int) (int, error) {
	if c > len(s) {
		return strconv.Atoi(s)
	}
	return 0, nil
}
