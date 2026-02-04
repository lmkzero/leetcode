package main

import "math"

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		sum := node.Val
		if left > 0 {
			sum += left
		}
		if right > 0 {
			sum += right
		}
		maxSum = max(maxSum, sum)
		if max(left, right) > 0 {
			return max(left, right) + node.Val
		}
		return node.Val
	}
	dfs(root)
	return maxSum
}
