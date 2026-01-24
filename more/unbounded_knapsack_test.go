package main

import "testing"

func Test_unboundedKnapsackDP(t *testing.T) {
	type args struct {
		wgt []int
		val []int
		cap int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unboundedKnapsackDP(tt.args.wgt, tt.args.val, tt.args.cap); got != tt.want {
				t.Errorf("unboundedKnapsackDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
