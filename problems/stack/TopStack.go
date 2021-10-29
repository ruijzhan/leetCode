package stack

import "errors"

//一道面试题： 实现一个Stack，能够查询当前堆栈中最大的元素。

type TopStack struct {
	stack []int
	tops  []int
}

var errEmptyStack = errors.New("empty stack")

func NewTopStack() *TopStack {
	return &TopStack{
		stack: make([]int, 0),
		tops:  make([]int, 0),
	}
}

func (s *TopStack) Top() (int, error) {
	if len(s.stack) == 0 {
		return 0, errEmptyStack
	}
	return s.tops[len(s.tops)-1], nil
}

func (s *TopStack) Push(n int) {
	s.stack = append(s.stack, n)
	if len(s.tops) == 0 || n >= s.tops[len(s.tops)-1] {
		s.tops = append(s.tops, n)
	}
}

func (s *TopStack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errEmptyStack
	}

	n := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	if n == s.tops[len(s.tops)-1] {
		s.tops = s.tops[:len(s.tops)-1]
	}
	return n, nil
}
