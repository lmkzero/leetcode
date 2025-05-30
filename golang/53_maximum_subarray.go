package main

func maxSubArray(nums []int) int {
	ans, f := nums[0], nums[0]
	for _, num := range nums[1:] {
		f = max(f, 0) + num
		ans = max(ans, f)
	}
	return ans
}
