package main

func subarraySum(nums []int, k int) int {
	cnt := map[int]int{0: 1}
	sum := 0
	ans := 0
	for _, num := range nums {
		sum += num
		ans += cnt[sum-k]
		cnt[sum]++
	}
	return ans
}
