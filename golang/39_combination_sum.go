package main

import (
	"slices"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var (
		comb []int
		ans  [][]int
	)
	sort.Ints(candidates)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, slices.Clone(comb))
			return
		}
		if target < candidates[idx] {
			return
		}
		for j := idx; j < len(candidates); j++ {
			comb = append(comb, candidates[j])
			dfs(target-candidates[j], j)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}
