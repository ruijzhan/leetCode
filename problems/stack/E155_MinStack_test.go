package stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 {
		t.Fatal()
	}
	s.Pop()
	if s.Top() != 0 {
		t.Fatal()
	}
	if s.GetMin() != -2 {
		t.Fatal()
	}
}
