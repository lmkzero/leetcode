package main

// 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil || left == right {
		return head
	}
	dummy := &ListNode{0, head}
	// 移动pre到left的前一个节点
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	// 反转left到right区间的节点
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
