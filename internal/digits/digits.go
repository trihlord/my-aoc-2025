package digits

import "strconv"

func IsRepeated(x int) bool {
	str := strconv.Itoa(x)
	l := len(str)
	for i, c := l/2, 0; i >= 1; i, c = i-1, 0 {
		for j := i; j <= l-i; j += i {
			if str[:i] == str[j:j+i] {
				c++
			}
		}
		if c == l/i-1 && l%i == 0 {
			return true
		}
	}
	return false
}

func IsRepeatedOnce(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	return s[l/2:] == s[:l/2]
}
