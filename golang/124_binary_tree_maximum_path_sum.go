package main

func maxPathSum(root *TreeNode) int {
	maxSum := 0
	if root == nil {
		return maxSum
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		maxSum = max(maxSum, left+right+node.Val)
		return max(left, right) + node.Val
	}
	dfs(root)
	return maxSum
}
