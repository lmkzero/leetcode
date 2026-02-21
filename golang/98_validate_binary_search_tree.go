package main

func isValidBST(root *TreeNode) bool {
	var (
		pre *TreeNode
		dfs func(*TreeNode) bool
	)
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if !dfs(root.Left) {
			return false
		}
		if pre != nil && pre.Val >= root.Val {
			return false
		}
		pre = root
		return dfs(root.Right)
	}
	return dfs(root)
}
