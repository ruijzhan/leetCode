package stack

import (
	"testing"
)

func TestTopStack(t *testing.T) {
	s := NewTopStack()
	if _, err := s.Top(); err == nil {
		t.Fatal()
	}
	if _, err := s.Pop(); err == nil {
		t.Fatal()
	}
	s.Push(5) // stack: [5]  top:5
	if top, err := s.Top(); err != nil || top != 5 {
		t.Fatal()
	}
	s.Push(4) // [5,4] 5
	if top, err := s.Top(); err != nil || top != 5 {
		t.Fatal()
	}
	s.Push(6) // [5,4,6] 6
	if top, err := s.Top(); err != nil || top != 6 {
		t.Fatal()
	}
	s.Push(5) // [5,4,6,5] 6
	if top, err := s.Top(); err != nil || top != 6 {
		t.Fatal()
	}
	if n, err := s.Pop(); err != nil || n != 5 { //[5,4,6] 6
		t.Fatal()
	}
	if top, err := s.Top(); err != nil || top != 6 {
		t.Fatal()
	}
	if n, err := s.Pop(); err != nil || n != 6 { //[5,4] 5
		t.Fatal()
	}
	if top, err := s.Top(); err != nil || top != 5 {
		t.Fatal()
	}
	if n, err := s.Pop(); err != nil || n != 4 { //[5] 5
		t.Fatal()
	}
	if top, err := s.Top(); err != nil || top != 5 {
		t.Fatal()
	}
}
