package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	if root == nil {
		return maxDiameter
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		maxDiameter = max(maxDiameter, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return maxDiameter
}
