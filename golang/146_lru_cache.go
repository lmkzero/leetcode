package main

type LRUCache struct {
	size, capacity int
	cache          map[int]*DoublyListNode
	head, tail     *DoublyListNode
}

func NewLRUCache(capacity int) LRUCache {
	head, tail := &DoublyListNode{}, &DoublyListNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DoublyListNode, capacity),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	l.removeNode(node)
	l.moveToHead(node)
	return node.val
}

func (l *LRUCache) removeNode(node *DoublyListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DoublyListNode) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.cache[key]
	if ok {
		node.val = value
		l.removeNode(node)
		l.moveToHead(node)
		return
	}
	newNode := &DoublyListNode{key: key, val: value}
	l.cache[key] = newNode
	l.size++
	l.moveToHead(newNode)
	if l.size > l.capacity {
		node := l.tail.prev
		delete(l.cache, node.key)
		l.removeNode(node)
		l.size--
	}
}
