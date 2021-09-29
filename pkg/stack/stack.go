package stack

type genericStack struct {
	array []interface{}
}

func (s *genericStack) Empty() bool {
	return len(s.array) == 0
}

func (s *genericStack) Len() int {
	return len(s.array)
}

func (s *genericStack) Push(e interface{}) {
	s.array = append(s.array, e)
}

func (s *genericStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, ErrEmptyStack
	}

	e := s.array[s.Len()-1]
	s.array = s.array[:s.Len()-1]

	return e, nil
}
