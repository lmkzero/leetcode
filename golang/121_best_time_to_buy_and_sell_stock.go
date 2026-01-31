package main

func maxProfitDP(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// maxProfit 计算给定股票价格数组中能获得的最大利润
// 参数 prices 是股票每日价格数组
// 返回值为能获得的最大利润，如果无法获利则返回0
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	max := 0
	for slow, fast := 0, 1; fast < len(prices); {
		curProfit := prices[fast] - prices[slow]
		if curProfit > max {
			max = curProfit
		}
		if curProfit < 0 {
			slow = fast
		}
		fast++
	}
	return max
}
