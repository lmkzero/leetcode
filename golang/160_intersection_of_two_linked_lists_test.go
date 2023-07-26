package main

import (
	"reflect"
	"testing"
)

func makeIntersectionList(vals1, vals2 []int, skip1, skip2 int) (*ListNode, *ListNode, *ListNode) {
	nodes1 := makeList(vals1[:skip1])
	nodes2 := makeList(vals2[:skip2])
	inters := makeList(vals1[skip1:])
	nodes1[skip1-1].Next = inters[0]
	nodes2[skip2-1].Next = inters[0]
	return nodes1[0], nodes2[0], inters[0]
}

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	headA1, headB1, inter1 := makeIntersectionList(
		[]int{4, 1, 8, 4, 5},
		[]int{5, 6, 1, 8, 4, 5},
		2,
		3,
	)
	headA2, headB2, inter2 := makeIntersectionList(
		[]int{1, 9, 1, 2, 4},
		[]int{3, 2, 4},
		3,
		1,
	)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{nil, nil},
			want: nil,
		},
		{
			name: "2",
			args: args{headA1, headB1},
			want: inter1,
		},
		{
			name: "3",
			args: args{headA2, headB2},
			want: inter2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
