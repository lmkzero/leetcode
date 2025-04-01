package main

func trap(height []int) int {
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
