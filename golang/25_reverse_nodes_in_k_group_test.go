package main

import (
	"reflect"
	"testing"
)

func Test_myReverse(t *testing.T) {
	type args struct {
		head *ListNode
		tail *ListNode
	}
	tests := []struct {
		name  string
		args  args
		want  *ListNode
		want1 *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := myReverse(tt.args.head, tt.args.tail)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myReverse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("myReverse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
