package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	path := make([]int, len(nums))
	used := make([]bool, len(nums))
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for j, num := range nums {
			if !used[j] {
				used[j] = true
				path[i] = num
				dfs(i + 1)
				used[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
