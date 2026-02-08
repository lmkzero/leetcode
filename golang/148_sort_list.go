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
	dummy := &ListNode{}
	tail := dummy
	for ; left != nil && right != nil; tail = tail.Next {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
			continue
		}
		tail.Next = right
		right = right.Next
	}
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}
	return dummy.Next
}
