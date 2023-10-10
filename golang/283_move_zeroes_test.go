package main

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	var (
		nums1    = []int{0, 1, 0, 3, 12}
		targets1 = []int{1, 3, 12, 0, 0}
		nums2    = []int{0}
		targets2 = []int{0}
	)
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{nums1},
			want: targets1,
		},
		{
			name: "2",
			args: args{nums2},
			want: targets2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !isIntSliceEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
