package main

func isPalindromeLinkedList(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	// 快慢指针寻找中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	// reverse后半段链表
	left, right := head, reverseList(slow)
	// 回文检查
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
