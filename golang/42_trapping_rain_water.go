package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
			continue
		}
		ans += rightMax - height[right]
		right--
	}
	return ans
}

func trapDP(height []int) int {
	n := len(height)
	peakIdx := 0
	for i, h := range height {
		if h > height[peakIdx] {
			peakIdx = i
		}
	}
	water := 0
	for i, leftPeak := 0, 0; i < peakIdx; i++ {
		if height[i] > leftPeak {
			leftPeak = height[i]
			continue
		}
		water += leftPeak - height[i]
	}
	for i, rightPeak := n-1, 0; i > peakIdx; i-- {
		if height[i] > rightPeak {
			rightPeak = height[i]
			continue
		}
		water += rightPeak - height[i]
	}
	return water
}
