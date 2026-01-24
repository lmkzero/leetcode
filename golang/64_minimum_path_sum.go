package main

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i][j-1]+grid[i][j], dp[i-1][j]+grid[i][j])
		}
	}
	return dp[row-1][col-1]
}
