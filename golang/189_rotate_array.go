package main

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	k %= len(nums)
	reverseArray(nums)
	reverseArray(nums[:k])
	reverseArray(nums[k:])
}

func reverseArray(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
