package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	nodes1 := makeList([]int{1, 2, 3, 4, 5})
	nodes2 := makeList([]int{1, 2})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{nodes1[0]},
			want: nodes1[len(nodes1)-1],
		},
		{
			name: "2",
			args: args{nodes2[0]},
			want: nodes1[len(nodes2)-1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
