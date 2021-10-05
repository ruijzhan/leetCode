package queue

type Iterator struct {
}

func (i *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}

func (i *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iter     *Iterator
	_next    int
	_hasNext bool
}

func Constructor1(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:     iter,
		_hasNext: iter.hasNext(),
		_next:    iter.next(),
	}
}

func (p *PeekingIterator) hasNext() bool {
	return p._hasNext
}

func (p *PeekingIterator) next() int {
	next := p._next
	p._hasNext = p.iter.hasNext()
	if p._hasNext {
		p._next = p.iter.next()
	}
	return next
}

func (p *PeekingIterator) peek() int {
	return p._next
}
