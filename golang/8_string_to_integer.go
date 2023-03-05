package main

import "math"

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		i        = 0
		negative = false
		res      = 0
	)
	// 去除前部字符
	for ; i < len(s) && s[i] == ' '; i++ {
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
	}
	// 确定数字
	for ; i < len(s); i++ {
		if s[i] < 48 || s[i] > 57 {
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
