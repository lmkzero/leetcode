package main

import "testing"

func Test_longestTaskChain(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tasks [][]int
		want  int
	}{
		{
			name:  "1",
			tasks: [][]int{{1, 5}, {2, 6}, {3, 8}, {4, 10}, {8, 9}, {5, 6}, {8, 10}, {6, 7}},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestTaskChain(tt.tasks)
			if got != tt.want {
				t.Errorf("longestTaskChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
