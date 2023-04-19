package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	ans := n
	for l <= r {
		mid := (r-l)>>1 + l
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
