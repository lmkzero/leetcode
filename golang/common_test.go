package main

import (
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
