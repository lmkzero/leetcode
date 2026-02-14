package main

import (
	"reflect"
	"testing"
)

func TestNewMinStack(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMinStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMinStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Push(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MinStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			this.Push(tt.args.val)
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			this.Pop()
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			if got := this.Top(); got != tt.want {
				t.Errorf("MinStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_GetMin(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("MinStack.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
