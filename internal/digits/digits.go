package digits

import "strconv"

func IsRepeated(x int) bool {
	str := strconv.Itoa(x)
	l := len(str)
	var sub string
	for i := l / 2; i >= 1; i-- {
		for j := 0; j < l-i; j++ {
			if j > 0 {
				//
			}
			sub = str[j : j+i]
		}
	}
	return false
}
