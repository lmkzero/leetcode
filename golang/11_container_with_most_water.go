package main

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var (
		left, right = 0, len(height) - 1
		mostWater   int
	)
	for left < right {
		curWater := (right - left) * min(height[left], height[right])
		mostWater = max(curWater, mostWater)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return mostWater
}
