package main

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	ans := make([][]int, 0, len(intervals))
	begin, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals[1:] {
		if end < interval[0] {
			ans = append(ans, []int{begin, end})
			begin, end = interval[0], interval[1]
			continue
		}
		if end < interval[1] {
			end = interval[1]
			continue
		}
	}
	ans = append(ans, []int{begin, end})
	return ans
}
