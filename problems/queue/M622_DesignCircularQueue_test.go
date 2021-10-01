package queue

import "testing"

func TestMyCircularQueue_AllFuncs(t *testing.T) {
	/*
		MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
		circularQueue.enQueue(1);  // 返回 true
		circularQueue.enQueue(2);  // 返回 true
		circularQueue.enQueue(3);  // 返回 true
		circularQueue.enQueue(4);  // 返回 false，队列已满
		circularQueue.Rear();  // 返回 3
		circularQueue.isFull();  // 返回 true
		circularQueue.deQueue();  // 返回 true
		circularQueue.enQueue(4);  // 返回 true
		circularQueue.Rear();  // 返回 4
	*/

	q := Constructor(3)
	if q.EnQueue(1) != true {
		t.Fatal()
	}
	if q.EnQueue(2) != true {
		t.Fatal()
	}
	if q.EnQueue(3) != true {
		t.Fatal()
	}
	if q.EnQueue(4) != false {
		t.Fatal()
	}
	if q.Rear() != 3 {
		t.Fatal()
	}
	if q.IsFull() != true {
		t.Fatal()
	}
	if q.DeQueue() != true {
		t.Fatal()
	}
	if q.EnQueue(4) != true {
		t.Fatal()
	}
	if q.Rear() != 4 {
		t.Fatal()
	}
}
