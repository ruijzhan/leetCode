package offer

type CQueue struct {
	s1 []int
	s2 []int
}

func NewCQueue() CQueue {
	return CQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.s2 = append(q.s2, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.s1) == 0 && len(q.s2) == 0 {
		return -1
	}
	if len(q.s1) > 0 {
		v := q.s1[len(q.s1)-1]
		q.s1 = q.s1[:len(q.s1)-1]
		return v
	}
	for len(q.s2) > 0 {
		v := q.s2[len(q.s2)-1]
		q.s2 = q.s2[:len(q.s2)-1]
		q.s1 = append(q.s1, v)
	}
	v := q.s1[len(q.s1)-1]
	q.s1 = q.s1[:len(q.s1)-1]
	return v
}
