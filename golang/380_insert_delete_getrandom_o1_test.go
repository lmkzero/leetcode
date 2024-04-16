package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want RandomizedSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Insert(t *testing.T) {
	type fields struct {
		nums    []int
		indices map[int]int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RandomizedSet{
				nums:    tt.fields.nums,
				indices: tt.fields.indices,
			}
			if got := rs.Insert(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Remove(t *testing.T) {
	type fields struct {
		nums    []int
		indices map[int]int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &RandomizedSet{
				nums:    tt.fields.nums,
				indices: tt.fields.indices,
			}
			if got := rs.Remove(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_GetRandom(t *testing.T) {
	type fields struct {
		nums    []int
		indices map[int]int
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
			rs := &RandomizedSet{
				nums:    tt.fields.nums,
				indices: tt.fields.indices,
			}
			if got := rs.GetRandom(); got != tt.want {
				t.Errorf("RandomizedSet.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
