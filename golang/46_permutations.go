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
		for j, u := range used {
			if !u {
				path[i] = nums[j]
				used[j] = true
				dfs(i + 1)
				used[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
