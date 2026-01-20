package main

import (
	"container/heap"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	hp := minHeap{}
	for _, num := range nums {
		heap.Push(&hp, num)
		if hp.Len() > k {
			heap.Pop(&hp)
		}
	}
	return hp.IntSlice[0]
}

type minHeap struct{ sort.IntSlice }

func (h minHeap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func (h *minHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *minHeap) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
