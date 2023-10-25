package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leafs := []*TreeNode{root}
	depth := 1
	for len(leafs) > 0 {
		row := len(leafs)
		for i := 0; i < row; i++ {
			cur := leafs[0]
			leafs = leafs[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				leafs = append(leafs, cur.Left)
			}
			if cur.Right != nil {
				leafs = append(leafs, cur.Right)
			}
		}
		depth++
	}
	return depth
}
