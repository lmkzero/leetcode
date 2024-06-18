package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	answers := make([]int, length)
	answers[0] = 1
	for i := 1; i < length; i++ {
		answers[i] = answers[i-1] * nums[i-1]
	}
	right := 1
	for i := length - 1; i >= 0; i-- {
		answers[i] = answers[i] * right
		right *= nums[i]
	}
	return answers
}
