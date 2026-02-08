package main

import (
	"reflect"
	"testing"
)

func TestNewMyQueue(t *testing.T) {
	tests := []struct {
		name string
		want MyQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMyQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Push(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			this.Push(tt.args.x)
		})
	}
}

func TestMyQueue_Pop(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			if got := this.Pop(); got != tt.want {
				t.Errorf("MyQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_move(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			this.move()
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			if got := this.Peek(); got != tt.want {
				t.Errorf("MyQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Empty(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyQueue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			if got := this.Empty(); got != tt.want {
				t.Errorf("MyQueue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
