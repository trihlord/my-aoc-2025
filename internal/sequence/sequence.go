package sequence

func Max(digits string, capacity int) string {
	stk := make([]byte, 0, capacity)
	dpth := len(digits) - capacity
	for i := 0; i < len(digits); i++ {
		char := digits[i]
		for len(stk) > 0 && char > stk[len(stk)-1] && dpth > 0 {
			stk = stk[:len(stk)-1]
			dpth--
		}
		if len(stk) < capacity {
			stk = append(stk, char)
			continue
		}
		dpth--
	}
	return string(stk)
}
