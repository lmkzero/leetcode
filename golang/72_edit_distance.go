package main

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dp := make([]int, m+1)
	// 状态转移：首行
	for j := 1; j <= m; j++ {
		dp[j] = j
	}
	// 状态转移：其余行
	for i := 1; i <= n; i++ {
		// 状态转移：首列
		leftUp := dp[0] // 暂存 dp[i-1, j-1]
		dp[0] = i
		// 状态转移：其余列
		for j := 1; j <= m; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				// 若两字符相等，则直接跳过此两字符
				dp[j] = leftUp
			} else {
				// 最少编辑步数 = 插入、删除、替换这三种操作的最少编辑步数 + 1
				dp[j] = min(min(dp[j-1], dp[j]), leftUp) + 1
			}
			leftUp = temp // 更新为下一轮的 dp[i-1, j-1]
		}
	}
	return dp[m]
}
