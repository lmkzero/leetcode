package main

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func makeList(vals []int) []*ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, 0, len(vals))
	for _, val := range vals {
		nodes = append(nodes, &ListNode{Val: val})
	}
	for i := 0; i < len(nodes); i++ {
		if i < len(nodes)-1 {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes
}

func makeLinkedList(vals []int) *ListNode {
	nodes := makeList(vals)
	if nodes == nil {
		return nil
	}
	return nodes[0]
}

func makeBinaryTree(vals []interface{}) []*TreeNode {
	if len(vals) == 0 {
		return nil
	}
	tree := make([]*TreeNode, 0, len(vals))
	for _, val := range vals {
		if intVal, ok := val.(int); ok {
			tree = append(tree, &TreeNode{Val: intVal})
		} else {
			tree = append(tree, &TreeNode{Val: math.MinInt})
		}
	}
	for idx, node := range tree {
		if node.Val == math.MinInt {
			continue
		}
		left := idx*2 + 1
		right := idx*2 + 2
		if left < len(tree) && tree[left].Val != math.MinInt {
			node.Left = tree[left]
		}
		if right < len(tree) && tree[right].Val != math.MinInt {
			node.Right = tree[right]
		}
	}
	return tree
}

func isIntSliceEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func isIntMatrixEqual(matrix1, matrix2 [][]int) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	for i, m1 := range matrix1 {
		if !isIntSliceEqual(m1, matrix2[i]) {
			return false
		}
	}
	return true
}

// matrix1, matrix2中的slice不重复
func isIntMatrixEqualWithoutOrder(matrix1, matrix2 [][]int) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	m1, m2 := make(map[string]struct{}, len(matrix1)), make(map[string]struct{}, len(matrix2))
	for i := 0; i < len(matrix1); i++ {
		m1[intSliceToString(matrix1[i])] = struct{}{}
		m2[intSliceToString(matrix2[i])] = struct{}{}
	}
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}
	return true
}

func intSliceToString(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	ss := make([]string, 0, len(nums))
	for _, num := range nums {
		ss = append(ss, strconv.FormatInt(int64(num), 10))
	}
	return strings.Join(ss, "-")
}

func Test_min(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1, 2},
			want: 1,
		},
		{
			name: "2",
			args: args{2, 3},
			want: 2,
		},
		{
			name: "3",
			args: args{0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{1, 2},
			want: 2,
		},
		{
			name: "2",
			args: args{2, 3},
			want: 3,
		},
		{
			name: "3",
			args: args{0, 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{0},
			want: 0,
		},
		{
			name: "2",
			args: args{-1},
			want: 1,
		},
		{
			name: "3",
			args: args{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
