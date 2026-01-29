package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = '0'                     // 访问过即置为0, 保证后序不重复搜索到
		if i-1 >= 0 && grid[i-1][j] == '1' { // 向左搜索
			dfs(i-1, j)
		}
		if i+1 < row && grid[i+1][j] == '1' { // 向右搜索
			dfs(i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' { // 向上搜索
			dfs(i, j-1)
		}
		if j+1 < col && grid[i][j+1] == '1' { // 向下搜索
			dfs(i, j+1)
		}
	}
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
