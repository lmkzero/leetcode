package main

import "testing"

func Test_isPalindromeLinkedList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := makeLinkedList([]int{1, 2, 2, 1})
	head2 := makeLinkedList([]int{1, 2})
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{head1},
			want: true,
		},
		{
			name: "2",
			args: args{head2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeLinkedList(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
