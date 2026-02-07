package main

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		ans = append(ans, queue[n-1].Val)
		for k := 0; k < n; k++ {
			node := queue[k]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return ans
}
