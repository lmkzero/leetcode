package main

import "slices"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	left := true
	ans := [][]int{}
	for len(queue) > 0 {
		t := []int{}
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			t = append(t, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		if !left {
			slices.Reverse(t)
		}
		ans = append(ans, t)
		left = !left
	}
	return ans
}
