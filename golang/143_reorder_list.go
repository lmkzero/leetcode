package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderListByIdx(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}
