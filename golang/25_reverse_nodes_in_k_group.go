package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		node := pre.Next
		next := tail.Next
		tail.Next = nil
		pre.Next = myReverse(node)
		node.Next = next
		pre = node
	}
	return dummy.Next
}

func myReverse(head *ListNode) *ListNode {
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
