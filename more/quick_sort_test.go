package main

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("quickSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_optimizedQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 2},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optimizedQuickSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("quickSort() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
