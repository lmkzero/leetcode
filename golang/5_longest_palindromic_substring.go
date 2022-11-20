package main

import "log"

// 1 <= s.length <= 1000
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	var (
		maxLen = 1
		begin  = 0
	)
	// init
	var dp [][]bool
	for i := 0; i < n; i++ {
		subDP := make([]bool, n)
		subDP[i] = true
		dp = append(dp, subDP)
	}
	log.Printf("%+v", dp)
	for i := 2; i <= n; i++ {
		for j := 0; j < n; j++ {
			index := i + j - 1
			if index >= n {
				break
			}
			if s[j] != s[index] {
				dp[j][index] = false
			} else {
				if index-j < 3 {
					dp[j][index] = true
				} else {
					dp[j][index] = dp[j+1][index-1]
				}
			}
			if dp[j][index] && index-j+1 > maxLen {
				maxLen = index - j + 1
				begin = j
			}
		}
	}
	return s[begin : begin+maxLen]
}
