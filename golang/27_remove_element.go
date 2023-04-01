package main

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	if right == 0 {
		return 0
	}
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
			continue
		}
		left++
	}
	return left
}
