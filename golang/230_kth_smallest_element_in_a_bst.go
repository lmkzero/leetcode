package main

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	count := 0
	ans := 0
	var dfs func(node *TreeNode, k int)
	dfs = func(node *TreeNode, k int) {
		if node == nil {
			return
		}
		dfs(node.Left, k)
		count++
		if count == k {
			ans = node.Val
			return
		}
		dfs(node.Right, k)
	}
	dfs(root, k)
	return ans
}
