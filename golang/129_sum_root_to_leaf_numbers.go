package main

func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(r *TreeNode, path int) int {
		if r == nil {
			return 0
		}
		path = path*10 + r.Val
		if r.Left == nil && r.Right == nil {
			return path
		}
		return dfs(r.Left, path) + dfs(r.Right, path)
	}
	return dfs(root, 0)
}
