package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	left, right := head, slow.Next
	slow.Next = nil
	left = sortList(left)
	right = sortList(right)
	return mergeTwoLists(left, right)
}
