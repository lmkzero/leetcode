package main

import "testing"

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tree1 := makeBinaryTree([]interface{}{3, 9, 20, "null", "null", 15, 7})
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{root: tree1[0]},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
