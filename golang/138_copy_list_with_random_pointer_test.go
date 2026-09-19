package main

import "testing"

func Test_copyRandomList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *Node
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList(tt.head)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
