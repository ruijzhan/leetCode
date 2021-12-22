package stack

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	s := Constructor2()
	s.Push(1)
	s.Push(2)
	if s.Top() != 2 {
		t.Fatal()
	}
	if s.Pop() != 2 {
		t.Fatal()
	}
	if s.Empty() {
		t.Fatal()
	}
}
