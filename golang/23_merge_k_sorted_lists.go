package main

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

func mergeKListsWithHeap(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{Val: -1}
	p := dummy
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

// PriorityQueue 优先队列
type PriorityQueue []*ListNode

// Len 优先队列长度
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less 优先队列元素比较
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

// Swap 优先队列交换方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 元素入队列
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

// Pop 元素出队列
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	node := old[len(old)-1]
	*pq = old[0 : len(old)-1]
	return node
}
