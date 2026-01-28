package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	res := [][]int{}
	for len(stack) > 0 {
		n := len(stack)
		level := make([]int, 0, n)
		for i := 0; i < n; i++ {
			node := stack[i]
			level = append(level, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[n:]
		res = append(res, level)
	}
	return res
}
