package main

// 1 <= left <= right <= n
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

// 反转链表前 N 个节点
var successor *ListNode // 后继节点

// 反转以 head 为起点的 n 个节点，返回新的头结点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// 以 head.Next 为起点，需要反转前 n - 1 个节点
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last
}
