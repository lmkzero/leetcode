package main

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
	// 初始化：长度为1的子串都是回文串
	var dp [][]bool
	for i := 0; i < n; i++ {
		subDP := make([]bool, n)
		subDP[i] = true
		dp = append(dp, subDP)
	}
	// 枚举子串长度len
	for len := 2; len <= n; len++ {
		// 枚举左边界left
		for left := 0; left < n; left++ {
			right := left + len - 1 // 右边界
			if right >= n {
				break
			}
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				if right-left < 3 {
					dp[left][right] = true
				} else {
					dp[left][right] = dp[left+1][right-1]
				}
			}
			// 记录最大长度
			if dp[left][right] && right-left+1 > maxLen {
				maxLen = right - left + 1
				begin = left
			}
		}
	}
	return s[begin : begin+maxLen]
}
