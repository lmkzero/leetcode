package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		path []int
		ans  [][]int
	)
	var dfs func(candidates []int, start int, target int)
	dfs = func(candidates []int, start int, target int) {
		if target == 0 { // target 不断减小，如果为0说明达到了目标值
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			// i != start 限制了这不对深度遍历到达的此值去重
			if i != start && candidates[i] == candidates[i-1] { // 去重
				continue
			}
			path = append(path, candidates[i])
			dfs(candidates, i+1, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	sort.Ints(candidates)
	dfs(candidates, 0, target)
	return ans
}
