package main

// flatten 将二叉树原地展开为单链表（LeetCode 114）。
// 展开后的顺序与先序遍历（前序）一致，且所有节点的 Left 均为 nil。
// 采用「寻找前驱节点」的迭代做法，整体时间复杂度 O(n)，空间复杂度 O(1)。
func flatten(root *TreeNode) {
	// 从根节点开始，逐个节点向右处理（根节点指针会被复用为遍历游标）
	for root != nil {
		if root.Left != nil {
			// 左子树不为空时，才需要把左子树搬到右边
			// 找到左子树中最右的节点（即先序遍历中 root 的直接后继的前驱）
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			// 把原先的右子树整棵接到 pre 的右边
			pre.Right = root.Right
			// 把左子树移到右边，成为新的右子树
			root.Right = root.Left
			// 断开左指针，保证左侧始终为空
			root.Left = nil
		}
		// 继续处理下一个节点（此时右子树的头部已就绪）
		root = root.Right
	}
}
