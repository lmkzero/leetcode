package main

import "testing"

func Test_maxDistance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arrays [][]int
		want   int
	}{
		{
			name:   "1",
			arrays: [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}},
			want:   4,
		},
		{
			name:   "2",
			arrays: [][]int{{1}, {1}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDistance(tt.arrays)
			if got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
