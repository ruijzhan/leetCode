package stack

import "errors"

var (
	ErrEmptyStack = errors.New("can not pop from empty stack")
)
