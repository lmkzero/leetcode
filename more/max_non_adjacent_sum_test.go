package main

import "testing"

func Test_maxNonAdjacentSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int64
		want int64
	}{
		{
			name: "1",
			nums: []int64{2, 6, 4, 1},
			want: 7,
		},
		{
			name: "2",
			nums: []int64{2, 7, 9, 3, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxNonAdjacentSum(tt.nums)
			if got != tt.want {
				t.Errorf("maxNonAdjacentSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
