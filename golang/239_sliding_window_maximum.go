package main

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
	hp := hp{}
	for i, v := range nums[:k-1] {
		heap.Push(&hp, pair{i, v})
	}
	ans := make([]int, 0, len(nums))
	for i := k - 1; i < len(nums); i++ {
		heap.Push(&hp, pair{i, nums[i]})
		for hp[0].i <= i-k {
			heap.Pop(&hp)
		}
		ans = append(ans, hp[0].v)
	}
	return ans
}

type pair struct {
	i, v int
}

type hp []pair

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.v > b.v || (a.v == b.v && i < j)
}

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(v any) { *h = append(*h, v.(pair)) }

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
