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

// optimizedQuickSort 优化版快速排序。
// 相比普通快速排序，这里做了一点优化：每次 partition 后，
// 只递归进入元素较少的那一侧，元素较多的一侧改为循环迭代处理。
// 这样可以将递归的最大深度控制在 O(log n)（递归深度由较小一侧决定），
// 从而避免在极端情况下（如几乎有序的数据）递归过深导致栈溢出。
func optimizedQuickSort(nums []int) {
	var sortFunc func(nums []int, left, right int)
	sortFunc = func(nums []int, left, right int) {
		// 使用迭代 + 递归混合的方式：大区间走循环，小区间走递归
		for left < right {
			// 以 nums[left] 为基准进行划分，返回基准最终所在位置 p
			p := partition(nums, left, right)
			// 比较基准左右两侧的长度，选择较短的区间递归，较长区间留在循环中处理
			if p-left < right-p {
				// 左侧较短：递归处理左区间 [left, p-1]
				sortFunc(nums, left, p-1)
				// 循环继续处理右区间 [p+1, right]
				left = p + 1
				continue
			}
			// 右侧较短（或相等）：递归处理右区间 [p+1, right]
			sortFunc(nums, p+1, right)
			// 循环继续处理左区间 [left, p-1]
			right = p - 1
		}
	}
	// 对整个数组排序
	sortFunc(nums, 0, len(nums)-1)
}
