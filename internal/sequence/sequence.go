package sequence

import "bytes"

func Max(digits string, capacity int) string {
	stk := []byte(digits)[:capacity:capacity]
Loop:
	for i := len(stk); i < len(digits); i++ {
		for j := 0; j < len(stk); j++ {
			otherStk := make([]byte, 0, len(stk))
			otherStk = append(otherStk, stk[:j]...)
			otherStk = append(otherStk, stk[j+1:]...)
			otherStk = append(otherStk, digits[i])
			if bytes.Compare(otherStk, stk) == 1 {
				stk = otherStk
				continue Loop
			}
		}
	}
	return string(stk)
}
