package sequence

import "strconv"

func Max(str string, cap int) (int, error) {
	if len(str) < cap {
		return strconv.Atoi(str)
	}
	return 0, nil
}
