package offer

import (
	"testing"
)

func TestCQueue(t *testing.T) {
	cq := NewCQueue()
	cq.AppendTail(3)
	if cq.DeleteHead() != 3 {
		t.Fatal()
	}
	if cq.DeleteHead() != -1 {
		t.Fatal()
	}
}

func TestCQueue2(t *testing.T) {
	cq := NewCQueue()
	if cq.DeleteHead() != -1 {
		t.Fatal()
	}
	cq.AppendTail(5)
	cq.AppendTail(2)
	if cq.DeleteHead() != 5 {
		t.Fatal()
	}
	if cq.DeleteHead() != 2 {
		t.Fatal()
	}
}
