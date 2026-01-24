package main

func changeCoinII(amount int, coins []int) int {
	n := len(coins)
	// 初始化 dp 表
	dp := make([]int, amount+1)
	dp[0] = 1
	// 状态转移
	for i := 1; i <= n; i++ {
		// 正序遍历
		for a := 1; a <= amount; a++ {
			if coins[i-1] > a {
				// 若超过目标金额，则不选硬币 i
				continue
			}
			// 不选和选硬币 i 这两种方案之和
			dp[a] = dp[a] + dp[a-coins[i-1]]
		}
	}
	return dp[amount]
}
