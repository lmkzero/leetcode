package main

func removeDuplicatesII(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	slow, fast := 2, 2
	for fast < length {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
