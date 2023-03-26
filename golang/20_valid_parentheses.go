package main

var pairs = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	stack := make([]byte, 0, len(s))
	for i := 0; i < n; i++ {
		if b, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != b {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
