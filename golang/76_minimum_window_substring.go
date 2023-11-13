package main

import "math"

func minWindow(s string, t string) string {
	chars := make(map[byte]int, len(t))
	window := make(map[byte]int, len(s))
	for i := 0; i < len(t); i++ {
		chars[t[i]] = chars[t[i]] + 1
	}
	left, right := 0, 0
	begin, length := 0, math.MaxUint32
	valid := 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := chars[c]; ok {
			window[c] = window[c] + 1
			if window[c] == chars[c] {
				valid++
			}
		}
		for valid == len(chars) {
			if length > right-left {
				begin = left
				length = right - left
			}
			moveC := s[left]
			if _, ok := chars[moveC]; !ok {
				window[moveC] = window[moveC] - 1
				left++
				continue
			}
			if window[moveC] == chars[moveC] {
				valid--
			}
			window[moveC] = window[moveC] - 1
			left++
		}
	}
	if length == math.MaxUint32 {
		return ""
	}
	return s[begin : begin+length]
}
