package linkedlist

type myNode struct {
	next *myNode
	val  int
}

type MyLinkedList struct {
	head *myNode
	len  int
}

func Constructor1() MyLinkedList {
	return MyLinkedList{
		head: new(myNode),
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index >= l.len {
		return -1
	}
	curr := l.head.next
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &myNode{val: val}
	l.len++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.len {
		return
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	n := &myNode{val: val, next: curr.next}
	curr.next = n
	l.len++
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index >= l.len {
		return
	}
	curr := l.head.next
	prev := l.head
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}
	prev.next = curr.next
	l.len--
}
