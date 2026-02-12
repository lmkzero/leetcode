package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	d := make(map[int]int)
	for i, v := range inorder {
		d[v] = i
	}
	var dfs func(int, int, int) *TreeNode
	dfs = func(i, j, n int) *TreeNode {
		if n <= 0 {
			return nil
		}
		v := preorder[i]
		idx := d[v]
		left := dfs(i+1, j, idx-j)
		right := dfs(i+1+idx-j, idx+1, n-1-(idx-j))
		return &TreeNode{v, left, right}
	}
	return dfs(0, 0, len(preorder))
}
