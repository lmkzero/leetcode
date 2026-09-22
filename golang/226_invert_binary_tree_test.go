package main

import "testing"

func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.root)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
