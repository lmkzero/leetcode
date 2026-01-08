package main

import (
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLRUCache(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRUCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {
	type fields struct {
		size     int
		capacity int
		cache    map[int]*DoublyListNode
		head     *DoublyListNode
		tail     *DoublyListNode
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LRUCache{
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
			if got := l.Get(tt.args.key); got != tt.want {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_removeNode(t *testing.T) {
	type fields struct {
		size     int
		capacity int
		cache    map[int]*DoublyListNode
		head     *DoublyListNode
		tail     *DoublyListNode
	}
	type args struct {
		node *DoublyListNode
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
			l := &LRUCache{
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
			l.removeNode(tt.args.node)
		})
	}
}

func TestLRUCache_moveToHead(t *testing.T) {
	type fields struct {
		size     int
		capacity int
		cache    map[int]*DoublyListNode
		head     *DoublyListNode
		tail     *DoublyListNode
	}
	type args struct {
		node *DoublyListNode
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
			l := &LRUCache{
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
			l.moveToHead(tt.args.node)
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		size     int
		capacity int
		cache    map[int]*DoublyListNode
		head     *DoublyListNode
		tail     *DoublyListNode
	}
	type args struct {
		key   int
		value int
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
			l := &LRUCache{
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
				cache:    tt.fields.cache,
				head:     tt.fields.head,
				tail:     tt.fields.tail,
			}
			l.Put(tt.args.key, tt.args.value)
		})
	}
}
