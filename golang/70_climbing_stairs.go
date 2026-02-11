package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func optimizedClimbStairs(n int) int {
	dp1, dp2 := 0, 1
	for i := 0; i < n; i++ {
		dp1, dp2 = dp2, dp1+dp2
	}
	return dp2
}
