package queue

type MyCircularQueue struct {
	q          []int
	head, tail int
	n          int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		q:    make([]int, k+1),
		head: 0,
		tail: 0,
		n:    k + 1,
	}
}

func (mq *MyCircularQueue) EnQueue(value int) bool {
	if mq.IsFull() {
		return false
	}
	mq.q[mq.tail] = value
	mq.tail = (mq.tail + 1) % mq.n

	return true
}

func (mq *MyCircularQueue) DeQueue() bool {
	if mq.IsEmpty() {
		return false
	}
	mq.head = (mq.head + 1) % mq.n

	return true
}

func (mq *MyCircularQueue) Front() int {
	if mq.IsEmpty() {
		return -1
	}

	return mq.q[mq.head]
}

func (mq *MyCircularQueue) Rear() int {
	if mq.IsEmpty() {
		return -1
	}
	if mq.tail == 0 {
		return mq.q[mq.n-1]
	}

	return mq.q[mq.tail-1]
}

func (mq *MyCircularQueue) IsEmpty() bool {
	return mq.head == mq.tail
}

func (mq *MyCircularQueue) IsFull() bool {
	return (mq.tail+1)%mq.n == mq.head
}
