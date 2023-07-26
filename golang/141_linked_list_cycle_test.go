package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1 := makeLinkedList([]int{1, 2, 3})
	head2, _ := makeCycleLinkedList(
		[]int{3, 2, 0, -4},
		1,
	)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{nil},
			want: false,
		},
		{
			name: "2",
			args: args{head1},
			want: false,
		},
		{
			name: "3",
			args: args{head2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
