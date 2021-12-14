package linkedlist

import (
	"testing"
)

func TestConstructor1(t *testing.T) {
	l := Constructor1()
	l.AddAtHead(1)
	l.DeleteAtIndex(0)
}

func TestConstructor2(t *testing.T) {
	l := Constructor1()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	if l.Get(1) != 2 {
		t.Fatal()
	}
	l.DeleteAtIndex(0)
	if l.Get(0) != 2 {
		t.Fatal()
	}
}

func TestConstructor3(t *testing.T) {
	l := Constructor1()
	l.AddAtHead(2)
	l.DeleteAtIndex(1)
	l.AddAtHead(2)
	l.AddAtHead(7)
	l.AddAtHead(3)
	l.AddAtHead(2)
	l.AddAtHead(5)
	l.AddAtTail(5)
	l.Get(5)
	l.DeleteAtIndex(6)
	l.DeleteAtIndex(4)
}

func TestConstructor4(t *testing.T) {
	l := Constructor1()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	if l.Get(1) != 2 {
		t.Fatal()
	}
	l.DeleteAtIndex(1)
	if l.Get(1) != 3 {
		t.Fatal()
	}
}

func TestConstructor5(t *testing.T) {
	l := Constructor1()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(3, 2)
}
