package main

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2 := m-1, n-1
	tail := m + n - 1
	for ; idx1 >= 0 && idx2 >= 0; tail-- {
		if nums1[idx1] >= nums2[idx2] {
			nums1[tail] = nums1[idx1]
			idx1--
		} else {
			nums1[tail] = nums2[idx2]
			idx2--
		}
	}
	if idx1 < 0 {
		for ; tail >= 0; tail-- {
			nums1[tail] = nums2[idx2]
			idx2--
		}
	}
}
