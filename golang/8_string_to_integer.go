package main

import "math"

func myAtoi(s string) int {
	var (
		i        = 0
		n        = len(s)
		negative = false
		res      = 0
	)
	if n == 0 {
		return 0
	}
	// 去除前部字符
	for i < n && s[i] == ' ' {
		i++
	}
	if i >= len(s) {
		return 0
	}
	// 确定正负
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		negative = true
	default:
	}
	// 确定数字
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		res = res*10 + int(s[i]-'0')
		if res > math.MaxInt32 && negative {
			return math.MinInt32
		}
		if res > math.MaxInt32 && !negative {
			return math.MaxInt32
		}
	}
	if negative {
		return -res
	}
	return res
}
