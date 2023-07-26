package main

import (
	"log"
	"reflect"
	"testing"
)

func makeCycleLinkedList(vals []int, pos int) (*ListNode, *ListNode) {
	nodes := makeList(vals)
	if nodes == nil {
		return nil, nil
	}
	if pos >= len(nodes) {
		log.Fatalf("illegal pos, len: %d, pos: %d", len(vals), pos)
		return nil, nil
	}
	nodes[len(nodes)-1].Next = nodes[pos]
	return nodes[0], nodes[pos]
}

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	head1, cycle1 := makeCycleLinkedList(
		[]int{3, 2, 0, -4},
		1,
	)
	head2, cycle2 := makeCycleLinkedList(
		[]int{1, 2},
		0,
	)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{nil},
			want: nil,
		},
		{
			name: "2",
			args: args{head1},
			want: cycle1,
		},
		{
			name: "3",
			args: args{head2},
			want: cycle2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
