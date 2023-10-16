package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	subsets := []int{}
	var dfs func(int, int, int)
	dfs = func(start, end, k int) {
		if len(subsets) == k {
			res = append(res, append([]int{}, subsets...))
			return
		}
		for i := start; i <= end; i++ {
			subsets = append(subsets, i)
			dfs(i+1, end, k)
			subsets = subsets[:len(subsets)-1]
		}
	}
	dfs(1, n, k)
	return res
}
