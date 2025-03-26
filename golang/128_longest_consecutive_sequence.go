package main

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	ss := map[int]int{}
	ans := 0
	for _, num := range nums {
		it := num
		for {
			if _, ok := m[it]; !ok {
				break
			}
			delete(m, it)
			it++
		}
		ss[num] = ss[it] + it - num
		ans = max(ans, ss[num])
	}
	return ans
}
