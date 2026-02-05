package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p := dummy
	for cur := head; cur != nil; {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if p.Next == cur {
			p = cur
			cur = cur.Next
			continue
		}
		p.Next = cur.Next
		cur = cur.Next
	}
	return dummy.Next
}
