package main

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	ans := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if grid[i][j] == 0 {
			return 0
		}
		ans := 1
		grid[i][j] = 0
		for k := 0; k < len(dirs); k++ {
			x, y := i+dirs[k][0], j+dirs[k][1]
			if x >= 0 && x < m && y >= 0 && y < n {
				ans += dfs(x, y)
			}
		}
		return ans
	}
	for r := range grid {
		for c := range grid[r] {
			ans = max(ans, dfs(r, c))
		}
	}
	return ans
}
