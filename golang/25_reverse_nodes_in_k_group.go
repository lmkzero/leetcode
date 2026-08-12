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
		pre.Next = reverseList(node)
		node.Next = next
		pre = node
	}
	return dummy.Next
}
