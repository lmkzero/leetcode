package main

func isBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := height(root.Left), height(root.Right)
		if l == -1 || r == -1 || abs(l-r) > 1 {
			return -1
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	return height(root) >= 0
}
