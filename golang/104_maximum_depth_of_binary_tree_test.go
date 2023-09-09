package main

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: makeBinaryTree([]interface{}{3, 9, 20, "null", "null", 15, 7})[0],
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: makeBinaryTree([]interface{}{1, "null", 2})[0],
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
