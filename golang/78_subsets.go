package main

func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}
	var dfs func(int)
	dfs = func(n int) {
		res = append(res, append([]int{}, subset...))
		for i := n; i < len(nums); i++ {
			subset = append(subset, nums[i])
			dfs(i + 1)
			subset = subset[:len(subset)-1]
		}
	}
	dfs(0)
	return res
}
