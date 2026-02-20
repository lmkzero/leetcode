package main

func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// 从第二天开始，如果当天股价大于前一天股价，则在前一天买入，当天卖出，即可获得利润。如果当天股价小于前一天股价，则不买入，不卖出。
// 也即是说，所有上涨交易日都做买卖，所有下跌交易日都不做买卖，最终获得的利润是最大的。
func maxProfitIIGreedy(prices []int) (ans int) {
	for i, v := range prices[1:] {
		t := v - prices[i]
		if t > 0 {
			ans += t
		}
	}
	return
}
