package main

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	fast, slow := 1, 1
	for fast < length {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
