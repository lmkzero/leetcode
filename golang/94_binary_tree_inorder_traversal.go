package main

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

func inorderTraversalWithStack(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, p.Val)
		p = p.Right
	}
	return ans
}
