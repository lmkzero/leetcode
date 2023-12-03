package main

import "testing"

func Test_maxProfitWithCooldown(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{1, 2, 3, 0, 2},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				prices: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitWithCooldown(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitWithCooldown() = %v, want %v", got, tt.want)
			}
		})
	}
}
