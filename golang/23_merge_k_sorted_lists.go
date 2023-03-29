package main

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}
