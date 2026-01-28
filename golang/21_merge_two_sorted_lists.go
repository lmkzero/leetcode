package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	prev := dummy
	for ; list1 != nil && list2 != nil; prev = prev.Next {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
			continue
		}
		prev.Next = list2
		list2 = list2.Next
	}
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return dummy.Next
}
