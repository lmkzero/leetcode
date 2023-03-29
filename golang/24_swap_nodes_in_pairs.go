package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummy.Next
}
