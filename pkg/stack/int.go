package stack

type stackInt struct {
	stack Interface
}

func NewIntStack() *stackInt {
	return &stackInt{
		stack: &genericStack{
			array: []interface{}{},
		},
	}
}

func (s *stackInt) Empty() bool {
	return s.stack.Empty()
}

func (s *stackInt) Len() int {
	return s.stack.Len()
}

func (s *stackInt) Push(i int) {
	s.stack.Push(i)
}

func (s *stackInt) Pop() (int, error) {
	v, err := s.stack.Pop()
	if err != nil {
		return 0, err
	}

	return v.(int), nil
}
