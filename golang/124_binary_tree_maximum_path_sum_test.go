package main

import "testing"

func Test_maxPathSum(t *testing.T) {
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
			args: args{root: makeBinaryTree([]interface{}{1, 2, 3})[0]},
			want: 6,
		},
		{
			name: "2",
			args: args{root: makeBinaryTree([]interface{}{-10, 9, 20, "null", "null", 15, 7})[0]},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
