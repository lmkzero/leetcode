package main

import "math"

func maxProfitIV(k int, prices []int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}
	if k > n/2 {
		return maxProfitII(prices)
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt32
	}
	for i := 0; i < n; i++ {
		for j := k; j > 0; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}
