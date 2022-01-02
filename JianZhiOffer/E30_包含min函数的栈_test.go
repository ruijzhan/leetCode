package offer

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.Min() != -3 {
		t.Fatal()
	}
	s.Pop()
	if s.Top() != 0 {
		t.Fatal()
	}
	if s.Min() != -2 {
		t.Fatal()
	}
}
