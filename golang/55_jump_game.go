package main

func canJump(nums []int) bool {
	rightMost := 0
	for i, n := range nums {
		if i <= rightMost {
			rightMost = max(rightMost, i+n)
			if rightMost >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
