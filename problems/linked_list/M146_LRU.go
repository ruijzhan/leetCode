package linkedlist

type dListNode struct {
	key  int
	val  int
	prev *dListNode
	next *dListNode
}

type LRUCache struct {
	cap      int
	mKeyNode map[int]*dListNode
	head     *dListNode
	tail     *dListNode
}

func Constructor(capacity int) LRUCache {
	head := new(dListNode)
	tail := new(dListNode)
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:      capacity,
		mKeyNode: make(map[int]*dListNode),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) mvToEnd(node *dListNode) {
	if node.next == lru.tail {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev

	lru.tail.prev.next = node
	node.prev = lru.tail.prev

	node.next = lru.tail
	lru.tail.prev = node
}

func (lru *LRUCache) add(key, value int) {
	node := &dListNode{
		key:  key,
		val:  value,
		prev: lru.tail.prev,
		next: lru.tail,
	}
	lru.tail.prev.next = node
	lru.tail.prev = node

	lru.mKeyNode[key] = node
}

func (lru *LRUCache) deleteOldest() {
	node := lru.head.next
	lru.head.next = node.next
	node.next.prev = lru.head

	delete(lru.mKeyNode, node.key)
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.mKeyNode[key]; ok {
		lru.mvToEnd(node)
		return node.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.mKeyNode[key]; ok {
		lru.mvToEnd(node)
		node.val = value
		return
	}

	if len(lru.mKeyNode) == lru.cap {
		lru.deleteOldest()
	}
	lru.add(key, value)
}
