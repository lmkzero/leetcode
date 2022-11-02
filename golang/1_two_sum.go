package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	m := make(map[int]int)
	for idx, v := range nums {
		if i, exist := m[target-v]; exist {
			res = append(res, i)
			res = append(res, idx)
			break
		}
		m[v] = idx
	}
	return res
}
