package main

import (
	"testing"
)

func Test_maxProfitII(t *testing.T) {
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
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitII(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitII() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitIIGreedy(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantAns: 7,
		},
		{
			name: "2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantAns: 4,
		},
		{
			name: "3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxProfitIIGreedy(tt.args.prices); gotAns != tt.wantAns {
				t.Errorf("maxProfitIIGreedy() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
