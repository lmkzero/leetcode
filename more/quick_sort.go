package main

// quickSort 快速排序
func quickSort(nums []int) {
	var sortFunc func(nums []int, left, right int)
	sortFunc = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		p := partition(nums, left, right)
		sortFunc(nums, left, p-1)
		sortFunc(nums, p+1, right)
	}
	sortFunc(nums, 0, len(nums)-1)
}

func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func optimizedQuickSort(nums []int) {
	var sortFunc func(nums []int, left, right int)
	sortFunc = func(nums []int, left, right int) {
		for left < right {
			p := partition(nums, left, right)
			if p-left < right-p {
				sortFunc(nums, left, p-1)
				left = p + 1
				continue
			}
			sortFunc(nums, p+1, right)
			right = p - 1
		}
	}
	sortFunc(nums, 0, len(nums)-1)
}
