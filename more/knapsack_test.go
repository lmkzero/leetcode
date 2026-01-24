package main

import "testing"

func Test_knapsackDP(t *testing.T) {
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
		{
			name: "1",
			args: args{
				wgt: []int{10, 20, 30, 40, 50},
				val: []int{50, 120, 150, 210, 240},
				cap: 50,
			},
			want: 270,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsackDP(tt.args.wgt, tt.args.val, tt.args.cap); got != tt.want {
				t.Errorf("knapsackDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
