package main

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = dummy.Next
		dummy.Next = cur
		cur = next
	}
	return dummy.Next
}
