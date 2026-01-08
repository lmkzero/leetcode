package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// DoublyListNode Definition for doubly-linked list.
type DoublyListNode struct {
	key, val   int
	next, prev *DoublyListNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
