package offer

type MinStack struct {
	data []int
	mins []int
}

func NewMinStack() MinStack {
	return MinStack{
		data: []int{},
		mins: []int{},
	}
}

func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)
	if len(s.mins) == 0 || x <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, x)
	}
}

func (s *MinStack) Pop() {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if v == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) Min() int {
	return s.mins[len(s.mins)-1]
}
